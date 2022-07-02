<template>
  <v-list>
    <v-expansion-panels>
      <v-expansion-panel v-for="(item, i) in items" :key="i" class="py-2">
        <method-mixin :user="user" :item="item" v-if="i === 0" />
        <method-telegram :user="user" :item="item" v-else-if="i === 1" />
        <method-signal :user="user" :item="item" v-else-if="i === 2" />
        <method-paid :user="user" :item="item" type="sms" v-else-if="i === 3" />
        <method-paid
          :item="item"
          :user="user"
          type="call"
          v-else-if="i === 4"
        />
      </v-expansion-panel>
    </v-expansion-panels>
  </v-list>
</template>

<script>
import MethodMixin from "~/components/MethodMixin.vue";
import MethodTelegram from "~/components/MethodTelegram.vue";
import MethodSignal from "~/components/MethodSignal.vue";
import MethodPaid from "~/components/MethodPaid.vue";

var images = {
  mixin: require("~/static/mixin.png"),
  telegram: require("~/static/telegram.png"),
  signal: require("~/static/signal.png"),
  sms: require("~/static/sms.jpg"),
  call: require("~/static/call.png"),
};

export default {
  components: {
    MethodMixin,
    MethodTelegram,
    MethodSignal,
    MethodPaid,
  },
  layout: "second",
  async fetch() {
    try {
      // fetch price
      if (this.$store.state.user.phone_number != undefined) {
        let number = this.$store.state.user.phone_number;
        let newpriceset = await this.$axios
          .get(`/price?number=${number}`)
          .then((response) => response.data);
        let priceset = JSON.stringify(newpriceset);
        this.$store.commit("updatePriceset", priceset);
      }

      // fetch all
      let token = this.$store.state.user.access_token;
      let config = {
        headers: {
          Authorization: `Bearer ${token}`,
          UserID: this.$store.state.user.user_id,
        },
      };

      const response = await this.$axios.get("/poll", config);
      if (response.status == 200) {
        let realdata = sortVault(response.data);
        this.$store.commit("poll", response.data);
        this.$store.commit("updateUser", realdata.user);
        this.$store.commit("updateVaults", realdata.vault);
      } else if (response.status == 401) {
        localStorage.removeItem("data");
        this.$router.push("/");
      }
    } catch (err) {
      console.log(err);
    }
  },
  fetchOnServer: false,
  computed: {
    user() {
      return this.$store.state.user;
    },
    items() {
      return [
        {
          id: 0,
          title: "Mixin Messenger",
          icon: images.mixin,
          value: this.user.identity_id,
          vaild: this.user.identity_id != "",
        },
        {
          id: 1,
          title: "Telegram",
          icon: images.telegram,
          value: this.user.tg_name,
          vaild: this.user.tg_id != "",
          content: this.$t("methods.telegram.name"),
        },
        {
          id: 2,
          title: "Signal",
          icon: images.signal,
          value: this.user.signal_number,
          vaild: this.user.signal_number != "",
          content: this.$t("methods.signal.number"),
        },
        {
          id: 3,
          title: this.$t("methods.sms"),
          icon: images.sms,
          value: this.user.phone_number,
          vaild: this.user.phone_number != "" && this.user.sms_balance != 0,
          content: this.$t("methods.phone.number"),
        },
        {
          id: 4,
          title: this.$t("methods.call"),
          icon: images.call,
          value: this.user.phone_number,
          vaild: this.user.phone_number != "" && this.user.call_balance != 0,
          content: this.$t("methods.phone.number"),
        },
      ];
    },
  },
};

function sortVault(data) {
  var vaults = [];
  if (data.vaults != null) {
    vaults = data.vaults;
    vaults.forEach(function (part, index, theArray) {
      theArray[index].alert_ratio = Math.round(theArray[index].alert_ratio);
    });
    var allvault = vaults.sort(function (a, b) {
      if (a.ratio == 0) return 1;
      if (b.ratio == 0) return -1;
      if (a.ratio < b.ratio) return -1;
      if (a.ratio > b.ratio) return 1;
    });
    data.vaults = allvault;
    return data;
  } else {
    return data;
  }
}
</script>