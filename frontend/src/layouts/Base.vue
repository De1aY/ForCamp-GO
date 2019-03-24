<template>
  <div class="layout">
    <header class="navbar"></header>
    <div class="menu"></div>
    <div class="content">
      <slot/>
    </div>
  </div>
</template>

<script>
    export default {
        name: "Base"
    }
</script>

<style lang="scss" scoped>
  
</style>
