<template>
  <div class="layout">
    <header class="navbar"></header>
    <div class="menu"></div>
    <div class="content">
      <router-view/>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Base',
};
</script>

<style lang="scss" scoped>
@import "../assets/scss/colors";

.layout {
  min-height: 100vh;
  display: grid;
  grid-template-rows: 80px auto;
  grid-template-columns: 80px auto;

  .menu {
    grid-row-start: 1;
    grid-row-end: span 2;
    grid-column-start: 1;
    border-right: 1.5px solid $neutralLighter;

    .menu-item {

    }

  }

  .content { background: $neutralLightest; }

}
</style>
