<template>
  <v-sheet dark class="overview rounded-lg">
    <v-row no-gutters class="py-4">
      <v-col cols="6">
        <div class="label-text mb-4">
          {{ $t("vault.overview.total") }}
        </div>

        <div class="total-text mb-2">
          <span>{{ total }}</span>
        </div>
      </v-col>

      <v-col cols="6">
        <div class="label-text mb-4">
          {{ $t("vault.overview.ratio") }}
        </div>

        <div class="total-text mb-2">
          <span>{{ averageRatio }}%</span>
        </div>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<script>
export default {
  props: ["vaults"],
  computed: {
    total(){
      return this.vaults.length
    },
    averageRatio(){
      return calcAvgRatio(this.vaults)
    }
  },
};

function calcAvgRatio(vaults) {
  let totalRatio = 0.0;
  let realLength = 0;
  try {
    for (var i = 0; i < vaults.length; i++) {
      if (!isNaN(vaults[i].ratio)) {
        totalRatio += parseFloat(vaults[i].ratio);
        realLength += 1;
      }
    }
    let result = (totalRatio / realLength).toFixed(2);
    if (isNaN(result)){
      return 0
    }
    return result
  
  } catch (e) {
    return 0;
  }
}
</script>


<style lang="scss" scoped>
.overview {
  text-align: center;
  position: sticky;
  height: 120px;
  background: linear-gradient(180deg, #272727 0%, #373737 100%) !important;
  .label-text {
    font-size: 12px;
    font-weight: 500;
    color: #808080;
  }
  .total-text {
    font-size: 28px;
    font-weight: 700;
  }
}
</style>