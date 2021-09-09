<template>
  <div class="pt-3 pb-2 mb-3 border-bottom">
    <router-link to="/customers/create" class="btn btn-sm btn-outline-secondary">Добавить клиента</router-link>
    <a href="javascript:void(0)" class="btn btn-sm btn-outline-info" @click="exportCSV">Экспорт списка клиентов</a>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">ФИО</th>
        <th scope="col">Пол</th>
        <th scope="col">Дата рождения</th>
        <th scope="col">Место рождения</th>
        <th scope="col">Изображение</th>
        <th scope="col">Действия</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="customer in customers" :key="customer.id">
        <td>{{ customer.id }}</td>
        <td>{{ customer.full_name }}</td>
        <td>{{ customer.gender }}</td>
        <td>{{ customer.birth_date }}</td>
        <td>{{ customer.birth_place }}</td>
        <td><img :src="customer.image" width="50" alt=""/></td>
        <td>
          <router-link :to="`/customers/${customer.id}/edit`" class="btn btn-sm btn-outline-secondary">Редактировать
          </router-link>
          <a href="javascript:void(0)" class="btn btn-sm btn-outline-secondary" @click="del(customer.id)">Удалить</a>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
  <Paginator :last-page="lastPage" @page-changed="load($event)"/>
</template>

<script lang="ts">
import {onMounted, ref} from "vue";
import axios from "axios";
import {Customer} from "@/models/customer";
import Paginator from "@/components/Paginator.vue";

export default {
  name: 'Customers',
  components: {Paginator},
  setup() {
    const customers = ref([]);
    const lastPage = ref(0);
    const load = async (page = 1) => {
      const {data} = await axios.get(`customers?page=${page}`)
      customers.value = data.data
      lastPage.value = data.meta.last_page
    }
    onMounted(load)

    const del = async (id: number) => {
      if (confirm('Вы уверены, что хотите удалить роль?')) {
        await axios.delete(`customers/${id}`)
        customers.value = customers.value.filter((c: Customer) => c.id !== id)
      }
    }

    const exportCSV = async () => {
      const {data} = await axios.post('export-customers', {}, {responseType: 'blob'});
      const blob = new Blob([data], {type: 'text/csv'});
      const  link = document.createElement('a');
      link.href = window.URL.createObjectURL(data);
      link.download = 'customers.csv';
      link.click();

    }

    return {
      customers,
      lastPage,
      del,
      load,
      exportCSV
    }
  }
}
</script>
