<template>
  <v-container class="fill-height pa-4">
    <v-layout align-center justify-center column>
      <div class="greyscale_1--text text-center body-2 mb-8">
        {{ $t("connect.text") }}
      </div>
      <v-btn
        rounded
        x-large
        depressed
        color="black"
        class="mb-10 white--text font-weight-medium"
        @click="toOauth()"
      >
        {{ $t("connect.wallet") }}
      </v-btn>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  layout: "DefaultLayout",
  name: "LoginPage",
  methods: {
    toOauth() {
      const auth = {
        phone: true,
        profile: true,
        contacts: false,
        assets: false,
        snapshots: false,
      };
      this.$bridge.login(auth, {
        redirect_url: process.env.oauth_url,
        client_id: process.env.client_id,
        code_challenge: false,
      });
    },
  },
  created(){
    this.$i18n.setLocale(this.$store.state.lang);
  }
};
</script>