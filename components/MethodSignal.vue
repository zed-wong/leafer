<template>
  <!-- Signal -->
  <div>
    <method-header :item="item" :user="user" />
    <v-expansion-panel-content>
      <v-row no-gutters>
        <!-- have number -->
        <template v-if="user.signal_number != ''">
          <v-col cols="12" class="text-center py-6 font-weight-medium">
            {{ item.content }}: {{ user.signal_number }}
          </v-col>
          <v-col cols="12" class="text-right pa-0">
            <v-spacer />
            <v-dialog v-model="dialog">
              <template v-slot:activator="{ on, attrs }">
                <v-btn text color="primary" v-on="on" v-bind="attrs">
                  {{ $t("methods.signal.update") }}
                </v-btn>
              </template>
              <action-update type="signal" :user="user" v-on:close-dialog="over"/>
            </v-dialog>
          </v-col>
        </template>

        <!-- set number -->
        <div v-else>
          <v-col class="py-2 px-1 text-subtitle-1 font-weight-medium">
            <span> {{ $t("methods.signal.text") }}: </span>
          </v-col>
          <v-col class="pb-6 pt-4 px-0 font-weight-medium">
            <MazPhoneNumberInput
              v-model="phone"
              @update="updateData"
              :preferredCountries="preferCountries"
              default-country-code="US"
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
import MethodHeader from "./MethodHeader.vue";
import { MazPhoneNumberInput } from "maz-ui";
export default {
  components: {
    MazPhoneNumberInput,
    MethodHeader,
  },
  props: ["item", "user"],
  data() {
    return {
      inputdata: {},
      phone: 0,
      dialog: false,
      preferCountries: [
        "CA",
        "US",
        "HK",
        "GB",
        "AE",
        "DE",
        "FR",
        "MY",
        "SG",
        "JP",
        "TH",
        "KR",
        "MO",
        "CN",
      ],
    };
  },
  methods: {
    over(){
      this.dialog = false;
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
            action: "signal",
            number: number,
          },
          config
        );
        if (resp.status == 200) {
          let tp = "signal"
          this.$store.commit("updateNumber", {tp, number});
          this.$router.push("/methods");
        }
      } catch (err) {
        console.log(err);
      }
    },
  },
};
</script>