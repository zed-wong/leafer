<template>
  <div>
    <!-- Mixin -->
    <method-header :item="item" :user="user" />
    <v-expansion-panel-content>
      <v-row no-gutters>
        <v-col cols="12" class="text-center pt-3 font-weight-medium">
          {{ $t("methods.mixin.id") }} : {{ item.value }}
        </v-col>
      </v-row>
    </v-expansion-panel-content>
  </div>
</template>

<script>
import MethodHeader from "./MethodHeader.vue";
export default {
  components: { MethodHeader },
  props: ["item","user"],
};
</script>