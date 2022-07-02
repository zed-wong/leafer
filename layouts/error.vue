<template>
  <div>
    <template v-if="error.statusCode === 404">
      <v-row justify="center" align="center" no-gutters>
        <v-col>
          <v-card color="transparent" flat class="text-center">
            <div class="text-h2 mt-5">404</div>
            <div class="display-1 mt-5">Page not found</div>
            <nuxt-link to="/">
              <div class="display-0 mt-10">Back</div>
            </nuxt-link>
          </v-card>
        </v-col>
      </v-row>
    </template>
    <template v-else>
      <h1>
        {{ otherError }}
      </h1>
      <NuxtLink to="/"> Home page </NuxtLink>
    </template>
  </div>
</template>

<script>
export default {
  layout: "second",
  props: {
    error: {
      type: Object,
      default: null,
    },
  },
  data() {
    return {
      pageNotFound: "404 Not Found",
      otherError: "An error occurred",
    };
  },
  head() {
    const title =
      this.error.statusCode === 404 ? this.pageNotFound : this.otherError;
    return {
      title,
    };
  },
};
</script>

<style scoped>
h1 {
  font-size: 20px;
}
.paragraph-text {
  font-size: 18px;
}
</style>
