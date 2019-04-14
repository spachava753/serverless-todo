<template>
  <div>
    List component
    <p>{{ simpleMsg }}</p>
    <div v-for="todo in todoList" :key="todo.Id">
      <p v-bind="todo.Title"></p>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ListComp",
  props: {},
  data() {
    return {
      simpleMsg: "Hello"
    };
  },
  computed: {
    todoList: function() {
      axios
        .get("https://pzsvbhabdl.execute-api.us-east-1.amazonaws.com/dev/list")
        .then(function(response) {
          console.log("Printing out response");
          console.log(response);
          todoList = response.data;
        })
        .catch(function(error) {
          console.log("Something happened and error occured while calling api");
          console.error(error);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
