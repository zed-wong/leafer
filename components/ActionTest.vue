<template>
  <v-card>
    <v-card-title class="text-h5 lighten-2">
      {{ title }}
    </v-card-title>

    <v-card-text class="font-weight-normal black--text pt-2">
      {{ text }}
    </v-card-text>

    <v-card-actions class="pt-0">
      <v-spacer></v-spacer>
      <v-btn color="primary" text @click="pay()">
        {{ $t("methods.test.pay") }}
      </v-btn>

      <v-snackbar v-model="snackbar" :timeout="snackTimeout">
        {{ $t("methods.test.snack") }}

        <template v-slot:action="{ attrs }">
          <v-btn color="pink" text v-bind="attrs" @click="snackbar = false">
            {{ $t("methods.test.snack.close") }}
          </v-btn>
        </template>
      </v-snackbar>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ["type", "user"],
  created() {
    if (this.type === "sms") {
      this.title = this.$t("methods.test.sms.title");
      this.text = this.$t("methods.test.sms.text");
    } else if (this.type === "call") {
      this.title = this.$t("methods.test.call.title");
      this.text = this.$t("methods.test.call.text");
    }
  },
  data() {
    return {
      title: "",
      text: "",
      snackbar: false,
      snackTimeout: 15000,
    };
  },
  methods: {
    pay() {
      let memo = {};
      switch (this.type) {
        case "call":
          memo = {
            action: "test",
            type: "call",
            price: "0.3",
            number: this.user.phone_number,
          };
          break;
        case "sms":
          memo = {
            action: "test",
            type: "sms",
            price: "0.1",
            number: this.user.phone_number,
          };
          break;
      }
      this.$bridge.payment({
        recipient: process.env.client_id,
        asset: process.env.pusdid,
        amount: memo.price,
        memo: JSON.stringify(memo),
      });
      this.snackbar = true
    },
  },
};
</script>