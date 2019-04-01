<template>
  <div class="categories">

  </div>
</template>

<script>
export default {
  name: 'Categories',
};
</script>

<style lang="scss" scoped>
  @import "../../../assets/scss/colors";

  .categories {
    margin: 60px;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: space-evenly;
    flex-wrap: wrap;

    > div { margin-bottom: 30px; }

    &-title {
      font-size: 24px;
      font-family: "Fira Sans Condensed", sans-serif;
      font-weight: 500;
    }
  }
</style>
