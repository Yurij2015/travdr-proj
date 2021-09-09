<template>
  <form @submit.prevent="submit">
    <h1 class="h3 mb-3 fw-normal">Редактировать пользователя</h1>
    <div class="mb-3">
      <label>Имя</label>
      <input v-model="data.first_name" name="first_name" class="form-control" placeholder="Имя">
    </div>
    <div class="mb-3">
      <label>Фамилия</label>
      <input v-model="data.last_name" name="last_name" class="form-control" placeholder="Фамилия">
    </div>
    <div class="mb-3">
      <label>Электронный адрес</label>
      <input v-model="data.email" name="email" type="email" class="form-control" placeholder="Электронный адрес">
    </div>
    <div class="mb-3">
      <label>Роль</label>
      <select v-model="data.role_id" name="role_id" class="form-control">
        <option v-for="role of roles" :key="role.id" :value="role.id">
          {{ role.name }}
        </option>
      </select>
    </div>
    <button class="w-100 btn btn-lg btn-primary">Сохранить</button>
  </form>
</template>

<script>
import {onMounted, reactive, ref} from "vue";
import axios from "axios";
import {useRoute, useRouter} from "vue-router";

export default {
  name: "UserEdit",
  setup() {
    const data = reactive({
      first_name: '',
      last_name: '',
      email: '',
      role_id: '',
    });
    const roles = ref([])
    const router = useRouter()
    const route = useRoute()
    onMounted(async () => {
      const rolesResponse = await axios.get('roles')
      roles.value = rolesResponse.data;

      const response = await  axios.get(`users/${route.params.id}`)
      data.first_name = response.data.first_name
      data.last_name = response.data.last_name
      data.email = response.data.email
      data.role_id = response.data.role.id
    })

    const submit = async () => {
      await axios.put(`users/${route.params.id}`, data)
      await router.push('/users')
    }
    return {
      data,
      submit,
      roles
    }
  }
}
</script>

<style scoped>

</style>