<template>
  <div class="pt-3 pb-2 mb-3 border-bottom">
    <router-link to="/roles/create" class="btn btn-sm btn-outline-secondary">Добавить роль</router-link>
  </div>
  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Наименование роли</th>
        <th scope="col">Действия</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="role in roles" :key="role.id">
        <td>{{ role.id }}</td>
        <td>{{ role.name }}</td>
        <td>
          <router-link :to="`/roles/${role.id}/edit`" class="btn btn-sm btn-outline-secondary">Редактировать
          </router-link>
          <a href="javascript:void(0)" class="btn btn-sm btn-outline-secondary" @click="del(role.id)">Удалить</a>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import {onMounted, ref} from "vue"
import axios from "axios"
import {Role} from "@/models/role";

export default {
  name: "Roles",
  setup() {
    const roles = ref([]);
    onMounted(async () => {
      const {data} = await axios.get('roles')
      roles.value = data
    })
    const del = async (id: number) => {
      if (confirm('Вы уверены, что хотите удалить роль?')) {
        await axios.delete(`roles/${id}`);
        roles.value = roles.value.filter((r: Role) => r.id !== id);
      }
    }
    return {
      roles,
      del
    }
  }
}
</script>

<style scoped>

</style>