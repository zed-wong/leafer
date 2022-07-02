<template>
  <v-card>
    <v-card-title class="text-h5">
      {{ $t("vault.renew.title") }}
    </v-card-title>
    <v-card-text class="py-2 font-weight-medium">
      {{ $t("vault.renew.text") }}
    </v-card-text>
    <v-card-text class="py-2 font-weight-medium">
      {{ $t("vault.renew.text2") }}
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn text @click="over">
        {{ $t("cancel") }}
      </v-btn>
      <v-btn color="primary" text @click="pay(vault.identity_id)">
        {{ $t("vault.detail.renew") }}
      </v-btn>
    </v-card-actions>

    <!-- overlay after payment -->
    <v-overlay :value="overlay">
      <v-row align="center" justify="center">
        <v-progress-circular
          indeterminate
          color="white"
          class="align-center mb-3"
        ></v-progress-circular>
      </v-row>
      <v-row>
        <v-col>
          <p class="text-center">{{ $t("methods.loading") }}</p>
        </v-col>
      </v-row>
      <v-row align="center" justify="center">
        <v-btn outlined rounded large class="mt-3" @click="overlay = false">
          {{ $t("cancel") }}
        </v-btn>
      </v-row>
    </v-overlay>
  </v-card>
</template>

<script>
export default {
  methods: {
    over: over,
    async pay(identityid) {
      this.overlay = true;

      let memo = {
        action: "renew",
        identityid: identityid,
        price: 12,
      };
      let p = this.$bridge.payment({
        recipient: process.env.client_id,
        asset: process.env.pusdid,
        amount: memo.price,
        memo: JSON.stringify(memo),
      });
      
      var initial = 0;
      var step = 5000;
      var timeout = 15000;
      while (initial < timeout) {
        await this.sleep(step);
        initial += step;
      }
      this.overlay = false;
      this.$router.push("/");
      this.$emit("close-dialog");
    },
    sleep(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
  },
  props: ["vault"],
  data() {
    return {
      overlay: false,
    };
  },
};

function over() {
  this.$emit("close-dialog");
}
</script>