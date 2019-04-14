<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col bg-dark text-white">
                <a class="navbar-brand">SPORTS STORE</a>
            </div>
        </div>
        <div class="row">
            <div class="col-3 bg-info p-2">
                <h4 class="text-white m-2">Categories</h4>
            </div>
            <div class="col-9 bg-success p-2">
                <h4 class="text-white m-2">Products</h4>
                  <div class="card-columns">

                    <div v-for="todo in todoList" :key="todo.Id" class="card m-1 p-1 bg-light">
                      <h4>{{todo.Title }}
                        <span class="badge badge-pill badge-primary float-right">
                          X
                        </span>
                      </h4>
                        <div class="card-text bg-white p-1">{{ todo.Description }}</div>
                    </div>
                  </div>

            </div>
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
      simpleMsg: "Hello",
      todoList: []
    };
  },
  created() {
    axios
      .get("https://pzsvbhabdl.execute-api.us-east-1.amazonaws.com/dev/list")
      .then(response => {
        this.todoList = response.data;
      })
      .catch(function(error) {
        console.log("Something happened and error occured while calling api");
        console.error(error);
      });
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
