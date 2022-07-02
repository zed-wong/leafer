<template>
  <v-card>
    <v-card-title class="text-h5 lighten-2">
      {{ title }}
    </v-card-title>

    <v-row no-gutters class="px-5 pt-3">
      <v-col
        class="d-flex justify-center"
        v-for="(item, i) in buyItems"
        :key="i"
      >
        <v-sheet
          rounded="lg"
          elevation="2"
          height="120"
          width="100"
          @click="select(i)"
          outlined
          :style="
            item.selected ? 'border: 1.5px solid #42A5F5' : 'border: none'
          "
        >
          <v-row no-gutters class="d-flex flex-column text-center fill-height">
            <v-col></v-col>
            <v-col class="justify-center">
              <span class="text-body-1">{{ item.num }}</span>
              <span>{{ unit }}</span>
            </v-col>
            <v-col></v-col>
            <v-col>
              <span class="text-caption"> $ </span>
              <span class="text-h6"> {{ item.price }} </span>
            </v-col>
            <v-col></v-col>
          </v-row>
        </v-sheet>
      </v-col>
    </v-row>

    <v-card-text> </v-card-text>

    <v-card-actions class="pt-0">
      <v-spacer></v-spacer>
      <v-btn color="primary" text @click="toPay">
        {{ $t("methods.buy") }}
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
  props: ["type", "user"],
  // BUG!
  created() {},
  computed: {
    priceset() {
      return JSON.parse(this.$store.state.priceset);
    },
    buyItems() {
      return this.fmtItems(this.type, this.buyitem, this.priceset)
    },
  },
  data() {
    return {
      title: "",
      unit: "",
      buyitem: [
        { id: 0, num: "", price: 0, selected: true },
        { id: 1, num: "", price: 0, selected: false },
      ],
      overlay: false,
    };
  },
  methods: {
    select(i) {
      if (i === 0) {
        this.buyitem[0].selected = true;
        this.buyitem[1].selected = false;
      } else if (i === 1) {
        this.buyitem[0].selected = false;
        this.buyitem[1].selected = true;
      }
    },
    async toPay() {
      this.overlay = !this.overlay;

      let memo = {};
      if (this.buyitem[0].selected) {
        memo = {
          action: "buy",
          type: this.type,
          plan: this.buyitem[0].id,
          num: this.buyitem[0].num,
          price: this.buyitem[0].price,
        };
      } else if (this.buyitem[1].selected) {
        memo = {
          action: "buy",
          type: this.type,
          plan: this.buyitem[1].id,
          num: this.buyitem[1].num,
          price: this.buyitem[1].price,
        };
      }

      this.$bridge.payment({
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
      this.$router.push("/methods");
    },

    sleep(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    fmtItems(type, item, priceset) {
      if (type === "sms") {
        this.title = this.$t("methods.buy.sms.title");
        this.unit = this.$t("methods.buy.sms.unit");
        item[0].num = priceset[0].num;
        item[0].price = priceset[0].price;
        item[1].num = priceset[1].num;
        item[1].price = priceset[1].price;
      } else if (type === "call") {
        this.title = this.$t("methods.buy.call.title");
        this.unit = this.$t("methods.buy.call.unit");
        item[0].num = priceset[2].num;
        item[0].price = priceset[2].price;
        item[1].num = priceset[3].num;
        item[1].price = priceset[3].price;
      }
      return item
    //       if (this.type === "sms") {
    //   this.title = this.$t("methods.buy.sms.title");
    //   this.unit = this.$t("methods.buy.sms.unit");
    //   this.buyitem[0].num = this.priceset[0].num;
    //   this.buyitem[0].price = this.priceset[0].price;
    //   this.buyitem[1].num = this.priceset[1].num;
    //   this.buyitem[1].price = this.priceset[1].price;
    // } else if (this.type === "call") {
    //   this.title = this.$t("methods.buy.call.title");
    //   this.unit = this.$t("methods.buy.call.unit");
    //   this.buyitem[0].num = this.priceset[2].num;
    //   this.buyitem[0].price = this.priceset[2].price;
    //   this.buyitem[1].num = this.priceset[3].num;
    //   this.buyitem[1].price = this.priceset[3].price;
    // }
    },
  },
};
</script>