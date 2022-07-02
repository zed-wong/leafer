<template>
  <div>
    <v-row no-gutters class="mx-5 mt-3">
      <span class="text-h4 font-weight-bold"> {{ $t("help.name") }} </span>
    </v-row>
    <v-row no-gutters class="mx-5 mt-8 d-flex flex-column">
      <v-col class="my-4" v-for="(q, i) in Questions" :key="i">
        <nuxt-link :to="localePath(`/questions/${i+1}`)">
          <span class="blue--text"> {{ q }} </span>
        </nuxt-link>
      </v-col>
    </v-row>
    <v-row no-gutters class="mx-5 mt-15 justify-center">
      <v-btn
        large
        rounded
        elevation="0"
        depressed
        @click="contactDev()"
        color="primary white--text"
      >
        {{ $t("contact.dev") }}
      </v-btn>
    </v-row>
  </div>
</template>

<script>
export default {
  layout: "questions",
  data() {
    return {
      Questions: [
        this.$t("help.1.title"),
        this.$t("help.2.title"),
        this.$t("help.3.title"),
        this.$t("help.4.title"),
        this.$t("help.5.title"),
        this.$t("help.6.title"),
      ],
    };
  },
  methods: {
    contactDev(){
      this.$bridge.popup('user', process.env.developer_id)
    }
  },
  created(){
    this.$i18n.setLocale(this.$store.state.lang);
  }
};
</script>

<style scoped>
a {
  text-decoration: none;
}
</style>