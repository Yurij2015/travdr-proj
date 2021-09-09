<template>
  <div class="pt-3 pb-2 mb-3 border-bottom">
    <router-link to="/agents/create" class="btn btn-sm btn-outline-secondary">Добавить агента</router-link>
    <a href="javascript:void(0)" class="btn btn-sm btn-outline-info" @click="exportCSV">Экспорт списка агентов</a>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">ФИО</th>
        <th scope="col">Номер телефона</th>
        <th scope="col">Организация</th>
        <th scope="col">Фотография</th>
        <th scope="col">Действия</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="agent in agents" :key="agent.id">
        <td>{{ agent.id }}</td>
        <td>
          <router-link :to="`/agents/${agent.id}/card`" class="text-decoration-none link-dark">{{ agent.full_name }}</router-link>
        </td>
        <td>{{ agent.phone_number }}</td>
        <td>{{ agent.organization.org_name }}</td>
        <td><img :src="agent.image" width="50" alt=""/></td>
        <td>
          <router-link :to="`/agents/${agent.id}/edit`" class="btn btn-sm btn-outline-secondary">Редактировать
          </router-link>
          <a href="javascript:void(0)" class="btn btn-sm btn-outline-secondary" @click="del(agent.id)">Удалить</a>
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
import {Agent} from "@/models/agent";
import Paginator from "@/components/Paginator.vue";

export default {
  name: 'Agents',
  components: {Paginator},
  setup() {
    const agents = ref([]);
    const lastPage = ref(0);
    const load = async (page = 1) => {
      const {data} = await axios.get(`agents?page=${page}`)
      agents.value = data.data
      lastPage.value = data.meta.last_page
    }
    onMounted(load)

    const del = async (id: number) => {
      if (confirm('Вы уверены, что хотите удалить агента?')) {
        await axios.delete(`agents/${id}`)
        agents.value = agents.value.filter((a: Agent) => a.id !== id)
      }
    }

    const exportCSV = async () => {
      const {data} = await axios.post('export-agents', {}, {responseType: 'blob'});
      const blob = new Blob([data], {type: 'text/csv'});
      const link = document.createElement('a');
      link.href = window.URL.createObjectURL(data);
      link.download = 'agents.csv';
      link.click();

    }
    return {
      agents: agents,
      lastPage,
      del,
      load,
      exportCSV
    }
  }
}
</script>
