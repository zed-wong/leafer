<template>
  <v-container class="fill-height">
    <v-row align="center" justify="center">
      <vue-loading
        type="bars"
        color="#20a0ff"
        :size="{ width: '40px', height: '50px' }"
      ></vue-loading>
    </v-row>
  </v-container>
</template>

<script>
import { VueLoading } from "vue-loading-template";
export default {
  methods: { getToken: getToken },
  created() {
    this.getToken();
  },
  components: {
    VueLoading,
  },
};

async function getToken() {
  try {
    if (process.client) {
      const code = this.$route.query.code;
      const conversationID = this.$bridge.conversationId || "";
      var lang = this.$bridge.getContext.locale || navigator.language;
      lang = lang.split("-")[0];
      if (code != undefined) {
        const params = new URLSearchParams({
          code: code,
          conversation_id: conversationID,
          lang: lang,
        });
        const response = await this.$axios.get("/mixinoauth", {
          params: params,
        });
        if (response.status == 200) {
          let realdata = sortVault(response.data);
          localStorage.setItem("data", JSON.stringify(realdata));
          this.$store.commit("updateUser", realdata.user);
          this.$store.commit("updateVaults", realdata.vault);
        }
        this.$router.push("/");
      }
    }
  } catch (error) {
    console.log(error);
  }
}

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