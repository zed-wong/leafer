<template>
  <v-app>
    <v-app-bar elevation="0" height="44px" color="white" fixed app>
      <v-row no-gutters>
        <v-col align-self="start">
          <v-btn icon @click="back">
            <v-icon> mdi-arrow-left </v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  methods: {
    back() {
      if (this.$route.name.substring(0, 4) == "help"){
        this.$router.push('/');
      }
      else {
        this.$router.go(-1);
      }
    },
  },
  created() {
    this.$i18n.setLocale(this.$store.state.lang);
  },
};
</script>