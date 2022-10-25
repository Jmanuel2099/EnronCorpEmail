<template>
  <div id="app">
    <Indexer />
  </div>
</template>

<script>
import Indexer from './components/Indexer.vue';

export default {
  name: 'App',
  components: {
    Indexer
}
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
