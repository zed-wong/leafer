<template>
  <!-- This is a layout for second layer pages (inside the avator popup sheet) -->
  <!-- v-container without fill-height -->
  <v-app>
    <v-app-bar elevation="0" height="44px" color="white" fixed app>
      <v-row no-gutters>
        <v-col align-self="start">
          <nuxt-link to="/">
            <v-btn icon>
              <v-icon> mdi-arrow-left </v-icon>
            </v-btn>
          </nuxt-link>
        </v-col>
        <v-col cols="4" align-self="center" class="d-flex justify-center">
          <nuxt-link to="/">
            <v-app-bar-title
              class="text-no-wrap font-weight-medium black--text"
            >
              <div>Leafer</div>
            </v-app-bar-title>
          </nuxt-link>
        </v-col>
        <v-col> </v-col>
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
      this.$router.go(-1);
    },
  },
  created() {
    this.$i18n.setLocale(this.$store.state.lang);
  },
};
</script>

<style scoped>
a {
  text-decoration: none;
}
</style>