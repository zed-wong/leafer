<template>
  <client-only>
    <v-app>
      <v-app-bar elevation="0" height="48px" :color="topbarColor" fixed app>
        <v-row no-gutters>
          <v-col align-self="start">
            <v-bottom-sheet v-model="sheet">
              <template v-slot:activator="{ on, attrs }">
                <v-avatar
                  height="32px"
                  min-width="32px"
                  width="32px"
                  v-bind="attrs"
                  v-on="on"
                >
                  <v-img :src="user.avatar" v-if="user.avatar" />
                  <v-img :src="unlogin" v-else />
                </v-avatar>
              </template>
              <popup-sheet :user="user" @off="off" />
            </v-bottom-sheet>
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
        <Nuxt />
      </v-main>
    </v-app>
  </client-only>
</template>

<script>
import PopupSheet from "../components/PopupSheet.vue";
export default {
  name: "DefaultLayout",
  components: PopupSheet,
  data() {
    return {
      topbarColor: "white",
      sheet: false,
      unlogin:
        "data:image/png;base64,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",
    };
  },
  methods: {
    off() {
      this.sheet = false;
    },
  },
  computed: {
    user() {
      return this.$store.state.user;
    },
  },
  created() {
    this.$i18n.setLocale(this.$store.state.lang);
  },
};
</script>

<style scoped>
.v-app-bar-title__content {
  width: 200px !important;
}
.v-sheet {
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
}
a {
  text-decoration: none;
}
</style>