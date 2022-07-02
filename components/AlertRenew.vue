<template>
  <v-dialog v-model="dialog">
    <template v-slot:activator="{ on, attrs }">
      <v-icon class="ml-2" color="red" v-bind="attrs" v-on="on">
        mdi-alert-octagon-outline
      </v-icon>
    </template>
    <v-card>
      <v-card-title class="text-h5">
        {{ $t("vault.renew.title") }}
      </v-card-title>
      <v-card-text class="py-2 font-weight-medium">
        {{ $t("vault.expired.text") }}
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn text @click="dialog=false">
          {{ $t("cancel") }}
        </v-btn>
        <v-btn color="primary" text @click="pay(vault.id)">
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
  </v-dialog>
</template>

<script>
export default {
  methods: {
    pay: pay,
  },
  props: ["vault"],
  data() {
    return {
      overlay: false,
      dialog: false,
    };
  },
};
function pay(vaultid) {
  this.overlay = true;

  let memo = {
    action: "renew",
    vaultid: vaultid,
    price: 12,
  };
  this.$bridge.payment({
    recipient: process.env.client_id,
    asset: process.env.pusdid,
    amount: memo.price,
    memo: JSON.stringify(memo),
  });
  //    update enddate in vault detail and localstorage
  //    this.$store.commit("updateDateByID", {id, date});
}
</script>