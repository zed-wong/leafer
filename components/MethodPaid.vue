<template>
  <div>
    <method-header :item="item" :user="user" />
    <v-expansion-panel-content>
      <v-row no-gutters>
        <!-- have number -->
        <template v-if="user.phone_number">
          <v-col cols="12" class="py-6 font-weight-medium">
            {{ item.content }}: {{ user.phone_number }}
          </v-col>

          <v-col cols="12" class="pt-1 pb-6 font-weight-medium">
            {{ $t("methods.left") }}: {{ remainBalance }}
          </v-col>

          <v-col cols="12" class="pa-0">
            <!-- Three buttons -->
            <v-row no-gutters>
              <v-col class="text-center">
                <v-dialog v-model="dialog0">
                  <template v-slot:activator="{ on0, attrs0 }">
                    <v-btn
                      text
                      color="primary"
                      @click="action('test')"
                      v-on="on0"
                      v-bind="attrs0"
                    >
                      {{ $t("methods.test") }}
                    </v-btn>
                  </template>
                  <action-test :type="type" :user="user" />
                </v-dialog>
              </v-col>

              <v-col class="text-center">
                <v-dialog v-model="dialog1">
                  <template v-slot:activator="{ on1, attrs1 }">
                    <v-btn
                      text
                      color="primary"
                      @click="action('buy')"
                      v-on="on1"
                      v-bind="attrs1"
                    >
                      {{ $t("methods.buy") }}
                    </v-btn>
                  </template>
                  <action-buy :type="type" :user="user" />
                </v-dialog>
              </v-col>

              <v-col class="text-center">
                <v-dialog v-model="dialog2">
                  <template v-slot:activator="{ on2, attrs2 }">
                    <v-btn
                      text
                      color="primary"
                      @click="action('update')"
                      v-on="on2"
                      v-bind="attrs2"
                    >
                      {{ $t("methods.update") }}
                    </v-btn>
                  </template>
                  <action-update
                    type="phone"
                    :user="user"
                    v-on:close-dialog="over"
                  />
                </v-dialog>
              </v-col>
            </v-row>
          </v-col>
        </template>

        <!-- set number -->
        <div v-else>
          <v-col class="py-2 px-1 text-subtitle-1 font-weight-medium">
            <span> {{ $t("methods.number.text") }}: </span>
          </v-col>
          <v-col class="pb-6 pt-4 px-0 font-weight-medium">
            <MazPhoneNumberInput
              v-model="phone"
              @update="updateData"
              default-country-code="CA"
              size="sm"
            />
          </v-col>
          <v-col cols="12" class="text-right pa-0">
            <v-spacer />
            <v-btn
              text
              color="primary"
              :disabled="!inputdata.isValid"
              @click="setNumber(phone)"
            >
              {{ $t("methods.set") }}
            </v-btn>
          </v-col>
        </div>
      </v-row>
    </v-expansion-panel-content>
  </div>
</template>

<script>
import { MazPhoneNumberInput } from "maz-ui";
import MethodHeader from "./MethodHeader.vue";
import ActionBuy from "./ActionBuy.vue";
import ActionTest from "./ActionTest.vue";
import ActionUpdate from "./ActionUpdate.vue";

export default {
  components: {
    MazPhoneNumberInput,
    MethodHeader,
    ActionBuy,
    ActionTest,
    ActionUpdate,
  },
  props: ["item", "type", "user"],
  computed: {
    remainBalance() {
      return this.type === "sms"
        ? this.user.sms_balance
        : this.user.call_balance;
    },
  },
  data() {
    return {
      inputdata: {},
      phone: 0,
      dialog0: false,
      dialog1: false,
      dialog2: false,
    };
  },
  methods: {
    over() {
      this.dialog2 = false;
    },
    updateData(data) {
      this.inputdata = data;
    },
    async setNumber(number) {
      let token = this.$store.state.user.access_token;
      let config = {
        headers: {
          Authorization: `Bearer ${token}`,
          UserID: this.$store.state.user.user_id,
        },
      };
      try {
        let resp = await this.$axios.put(
          "/update/number",
          {
            action: "phone",
            number: number,
          },
          config
        );
        if (resp.status == 200) {
          let tp = "phone";
          this.$store.commit("updateNumber", { tp, number });
          this.$router.push("/methods");
        }
      } catch (err) {
        console.log(err);
      }
    },
    action(type) {
      switch (type) {
        case "test":
          this.dialog0 = true;
          break;
        case "buy":
          this.dialog1 = true;
          break;
        case "update":
          this.dialog2 = true;
          break;
      }
    },
    off() {
      this.dialog = false;
    },
  },
};
</script>