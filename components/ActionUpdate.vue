<template>
  <v-card>
    <v-card-title class="text-h5 lighten-2">
      {{ $t("methods.update.title") }}
    </v-card-title>

    <v-card-text>
      <v-row no-gutters class="d-flex flex-column">
        <v-col class="py-2 px-1 text-subtitle-1 font-weight-medium">
          <span> {{ $t("methods.update.text") }}: </span>
        </v-col>
        <v-col class="pb-6 pt-4 px-0 font-weight-medium">
          <MazPhoneNumberInput
            v-model="phone"
            @update="updateData"
            :noCountrySelector="true"
            :placeholder="placeholder"
            default-country-code="CA"
            size="sm"
          />
        </v-col>
      </v-row>
    </v-card-text>

    <v-card-actions class="text-right pt-0">
      <v-spacer></v-spacer>
      <v-btn
        color="primary"
        text
        :disabled="!inputdata.isValid"
        @click="update(phone)"
      >
        {{ $t("methods.update") }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { MazPhoneNumberInput } from "maz-ui";
export default {
  components: {
    MazPhoneNumberInput,
  },
  props: ["user", "type"],
  data() {
    return {
      inputdata: {},
      phone: 0,
      placeholder: this.$t("methods.update.example"),
    };
  },
  methods: {
    updateData(data) {
      this.inputdata = data;
      if (data.isValid) {
        this.placeholder = this.$t("methods.update.valid");
      } else {
        this.placeholder = this.$t("methods.update.example");
      }
    },
    async update(number) {
      let token = this.$store.state.user.access_token;
      let config = {
        headers: {
          Authorization: `Bearer ${token}`,
          UserID: this.$store.state.user.user_id,
        },
      };
      this.$axios.put(
        "/update/number",
        {
          action: this.type,
          number: number,
        },
        config
      );
      let tp = this.type;
      this.$store.commit("updateNumber", { tp, number });
      this.$emit("close-dialog");

      if (this.type == "phone") {
        if (this.$store.state.user.phone_number != undefined) {
          let newpriceset = await this.$axios
            .get(`/price?number=${number}`)
            .then((response) => response.data);
          let priceset = JSON.stringify(newpriceset);
          this.$store.commit("updatePriceset", priceset);
          this.$router.push("/methods");
        }
      }
    },
  },
};
</script>