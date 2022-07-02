<template>
  <div>
    <v-row no-gutters class="mx-5 mt-3">
      <span class="text-h5 font-weight-bold"> {{ Answer.title }} </span>
    </v-row>
    <v-row no-gutters class="mx-5 mt-8 d-flex flex-column">
      <span> {{ Answer.article }} </span>
    </v-row>
  </div>
</template>

<script>
export default {
  layout: "questions",
  data() {
    return {
      Answer: {
        title: this.$t("help.4.title"),
        article: this.$t("help.4.article"),
      },
    };
  },
  created(){
    this.$i18n.setLocale(this.$store.state.lang);
  }
};
</script>