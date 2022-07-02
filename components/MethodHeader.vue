<template>
  <!-- Free -->
  <v-expansion-panel-header
    disable-icon-rotate
    class="pa-1 pr-3"
    v-if="item.id < 3"
  >
    <v-list-item>
      <v-list-item-avatar size="36px">
        <v-img :src="item.icon" />
      </v-list-item-avatar>
      <v-list-item-title class="font-weight-medium pl-1 text-subtitle-1">
        {{ item.title }}
      </v-list-item-title>
      <v-divider />
    </v-list-item>
    <template v-slot:actions>
      <v-icon color="green" v-if="item.vaild">
        mdi-checkbox-marked-circle
      </v-icon>
      <v-icon color="error" v-else> mdi-alert-circle </v-icon>
    </template>
  </v-expansion-panel-header>

  <!-- Charge -->
  <v-expansion-panel-header disable-icon-rotate class="pa-1 pr-3" v-else>
    <template v-slot:default="{ open }">
      <v-list-item>
        <v-list-item-avatar size="36px">
          <v-img :src="item.icon" />
        </v-list-item-avatar>
        <v-list-item-title class="font-weight-medium pl-1 text-subtitle-1">
          {{ item.title }}
        </v-list-item-title>
        <v-divider />
      </v-list-item>
      <v-fade-transition leave-absolute>
        <span v-if="!open && item.id==3" class="text-right mr-4 font-weight-light"> {{ user.sms_balance }} </span>
        <span v-if="!open && item.id==4" class="text-right mr-4 font-weight-light"> {{ user.call_balance }} </span>
      </v-fade-transition>
    </template>
    <template v-slot:actions>
      <v-icon color="green" v-if="item.vaild">
        mdi-checkbox-marked-circle
      </v-icon>
      <v-icon color="error" v-else> mdi-alert-circle </v-icon>
    </template>
  </v-expansion-panel-header>
</template>

<script>
export default {
  props: ["item", "user", "i"],
};
</script>