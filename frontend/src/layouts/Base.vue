<template>
  <div class="layout">
    <header class="navbar"></header>
    <div class="menu">
      <div class="menu-action">
        <font-awesome-icon icon="bars" class="menu-icon" />
      </div>
      <router-link to="profile" class="menu-item"
                   v-bind:class="{'menu-item--active': $router.currentRoute.name === 'profile'}">
        <font-awesome-icon icon="user-alt" class="menu-icon" />
      </router-link>
      <router-link to="statistics" class="menu-item"
                   v-bind:class="{'menu-item--active': $router.currentRoute.name === 'statistics'}">
          <font-awesome-icon icon="chart-bar" class="menu-icon" />
      </router-link>
    </div>
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
  grid-template-rows: 70px auto;
  grid-template-columns: 70px auto;

  .navbar {
    background: #fff;
  }

  .menu {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: stretch;
    grid-row-start: 1;
    grid-row-end: span 2;
    grid-column-start: 1;
    background: #fff;
    border-right: 1.5px solid $neutralLighter;

    &-action {
      height: 70px;
      display: flex;
      justify-content: center;
      align-items: center;

      &:hover {
        cursor: pointer;

        .menu-icon { color: $primary; }

      }

    }

    &-item {
      @extend .menu-action;
      border-right: 4px solid #fff;
      border-left: 4px solid #fff;
      transition: border 0.2s ease-in-out;

      &:hover { border-left: 4px solid $primary; }

      &--active {
        border-left: 4px solid $primary;

        .menu-icon { color: $primary; }

      }

    }

    &-icon {
      color: $neutralDarker;
      transition: color 0.2s ease-in-out;
    }

  }

  .content {
    padding: 60px 0 0 60px;
    background: $neutralLightest;
  }

}
</style>
