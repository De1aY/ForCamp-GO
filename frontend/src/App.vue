<template>
  <layout>
    <router-view/>
  </layout>
</template>

<script>
  import Layout from "@/layouts/Base";
  export default {
    components: {
      Layout
    }
  }
</script>

<style lang="scss">
  @import "assets/scss/colors";
</style>
