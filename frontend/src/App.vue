<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss">
@import "assets/scss/colors";
@import url('https://fonts.googleapis.com/css?family=
  Fira+Sans+Condensed:300,400,500,600,700|
  Montserrat:400,600,800|
  Open+Sans:300,400,700
  &subset=cyrillic'
);

body {
  margin: 0;
  padding: 0;
}

h1 { font-family: "Fira Sans Condensed", sans-serif; }
</style>
