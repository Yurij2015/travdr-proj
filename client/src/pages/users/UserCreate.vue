<template>
  <form @submit.prevent="submit">
    <h1 class="h3 mb-3 fw-normal">Добавить пользователя</h1>
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
import {useRouter} from "vue-router";

export default {
  name: "UserCreate",
  setup() {
    const data = reactive({
      first_name: '',
      last_name: '',
      email: '',
      role_id: '',
    });
    const roles = ref([])
    const router = useRouter()
    onMounted(async () => {
      const {data} = await axios.get('roles')
      roles.value = data;
    })

    const submit = async () => {
      await axios.post('users', data)
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