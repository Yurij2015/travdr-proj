<template>
  <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
    <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3" href="#">TravelersDreams</a>
<!--    <input class="form-control form-control-dark w-100" type="text" placeholder="Search" aria-label="Search">-->
    <div class="navbar-nav">
      <router-link class="nav-link px-3" href="#" to="/profile">{{ name }}</router-link>
      <router-link class="nav-link px-3" href="#" to="/login" @click="logout">Выход</router-link>
    </div>
  </header>
</template>

<script lang="ts">
import {computed, ref, watch} from "vue";
import axios from "axios";
import {useStore} from "vuex";


export default {
  name: "Nav",
  setup() {
    const name = ref('');
    const store = useStore();
    const user = computed(() => store.state.User.user);

    watch(user, () => {
      name.value = user.value.email;
    });

    const logout = async () => {
      await axios.post('logout')
    }
    return {
      name,
      logout
    }
  }
}
</script>