<template>
  <v-container class="fill-height pa-4">
    <novault v-if="vault == undefined || vault.length == 0" />
    <allvault v-else />
  </v-container>
</template>

<script>
import novault from "./novault.vue";
import allvault from "./allvault.vue";
export default {
  components: { novault, allvault },
  name: "index",
  computed: {
    vault() {
      return this.$store.state.vault;
    },
  },
  created() {
    if (process.client) {
      this.$i18n.setLocale(this.$store.state.lang);
      if (!localStorage.getItem("data")) {
        this.$router.push("/login");
      }
    }
  },
  async fetch() {
    try {
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
      }
    } catch (err) {
      console.log(err);
      localStorage.removeItem("data");
      this.$router.push("/login");
      this.$store.commit("updateVaults", []);
      this.$store.commit("updateUser", {});
    }
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

<style>
body {
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>