<template>
  <div class="pt-3 pb-2 mb-3 border-bottom">
    <router-link to="/contracts/create" class="btn btn-sm btn-outline-secondary">Добавить договор</router-link>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Номер договора</th>
        <th scope="col">Начало путешествия</th>
        <th scope="col">Окончание путешествия</th>
        <th scope="col">Организация</th>
        <th scope="col">Соглашение</th>
        <th scope="col">Агент</th>
        <th scope="col">Клиент</th>
        <th scope="col">Действия</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="contract in contracts" :key="contract.id">
        <td>{{ contract.id }}</td>
        <td>
          <router-link :to="`/contracts/${contract.id}/card`" class="text-decoration-none link-dark">{{
              contract.contract_number
            }}
          </router-link>
        </td>
        <td>{{ contract.start_trip_date }}</td>
        <td>{{ contract.end_trip_date }}</td>
        <td>{{ contract.organization.org_name }}</td>
        <td>{{ contract.agreement.agreement_name }}</td>
        <td>{{ contract.agent.full_name }}</td>
        <td>{{ contract.customer.full_name }}</td>
        <td>
          <router-link :to="`/contracts/${contract.id}/edit`" class="btn btn-sm btn-outline-secondary">Редактировать
          </router-link>
          <a href="javascript:void(0)" class="btn btn-sm btn-outline-secondary" @click="del(contract.id)">Удалить</a>
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
import {Contract} from "@/models/contract";
import Paginator from "@/components/Paginator.vue";

export default {
  name: 'Contracts',
  components: {Paginator},
  setup() {
    const contracts = ref([]);
    const lastPage = ref(0);
    const load = async (page = 1) => {
      const {data} = await axios.get(`contracts?page=${page}`)
      contracts.value = data.data
      lastPage.value = data.meta.last_page
    }
    onMounted(load)
    const del = async (id: number) => {
      if (confirm('Вы уверены, что хотите удалить договор?')) {
        await axios.delete(`contracts/${id}`)
        contracts.value = contracts.value.filter((a: Contract) => a.id !== id)
      }
    }
    return {
      contracts: contracts,
      lastPage,
      del,
      load,
    }
  }
}
</script>
