<template>
  <div class="pt-3 pb-2 mb-3 border-bottom">
    <router-link to="/agreements/create" class="btn btn-sm btn-outline-secondary">Добавить соглашение</router-link>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Номер соглашения</th>
        <th scope="col">Краткое описание</th>
        <th scope="col">Содержимое</th>
        <th scope="col">Действия</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="agreement in agreements" :key="agreement.id">
        <td>{{ agreement.id }}</td>
        <td>
          <router-link :to="`/agreements/${agreement.id}/card`" class="text-decoration-none link-dark">{{
              agreement.agreement_name
            }}
          </router-link>
        </td>
        <td>{{ agreement.description }}</td>
        <td v-html="agreement.content"/>
        <td>
          <router-link :to="`/agreements/${agreement.id}/edit`" class="btn btn-sm btn-outline-secondary">Редактировать
          </router-link>
          <a href="javascript:void(0)" class="btn btn-sm btn-outline-secondary" @click="del(agreement.id)">Удалить</a>
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
import {Agreement} from "@/models/agreement";
import Paginator from "@/components/Paginator.vue";

export default {
  name: 'Agreements',
  components: {Paginator},
  setup() {
    const agreements = ref([]);
    const lastPage = ref(0);
    const load = async (page = 1) => {
      const {data} = await axios.get(`agreements?page=${page}`)
      agreements.value = data.data
      lastPage.value = data.meta.last_page
    }
    onMounted(load)
    const del = async (id: number) => {
      if (confirm('Вы уверены, что хотите удалить соглашение?')) {
        await axios.delete(`agreements/${id}`)
        agreements.value = agreements.value.filter((a: Agreement) => a.id !== id)
      }
    }
    return {
      agreements: agreements,
      lastPage,
      del,
      load,
    }
  }
}
</script>
