<template>
  <div class="row">
    <div class="col-md-12 mt-3">
      <form @submit.prevent="submit">
        <h1 class="h3 mb-3 fw-normal">Добавить договор</h1>
        <div class="mb-3">
          <label for="contract_number">Номер договора</label>
          <input name="contract_number" id="contract_number" class="form-control" placeholder="Номер договора" required v-model="data.contract_number">
        </div>
        <div class="mb-3">
          <label for="start_trip_date">Начало поездки</label>
          <input name="start_trip_date" id="start_trip_date" type="date" class="form-control"
                 v-model="data.start_trip_date" required>
        </div>
        <div class="mb-3">
          <label for="end_trip_date">Окончание поездки</label>
          <input name="end_trip_date" id="end_trip_date" type="date" class="form-control"
                 v-model="data.end_trip_date" required>
        </div>
        <div class="mb-3">
          <label for="organization_id">Организация</label>
          <select v-model="data.organization_id" id="organization_id" name="organization_id" class="form-control" required>
            <option v-for="organization of organizations" :key="organization.id" :value="organization.id">
              {{ organization.org_name }}
            </option>
          </select>
        </div>
        <div class="mb-3">
          <label for="agreement_id">Соглашение</label>
          <select v-model="data.agreement_id" id="agreement_id" name="agreement_id" class="form-control" required>
            <option v-for="agreement of agreements" :key="agreement.id" :value="agreement.id">
              {{ agreement.agreement_name }}
            </option>
          </select>
        </div>
        <div class="mb-3">
          <label for="agent_id">Агент</label>
          <select v-model="data.agent_id" id="agent_id" name="agent_id" class="form-control" required>
            <option v-for="agent of agents" :key="agent.id" :value="agent.id">
              {{ agent.full_name }}
            </option>
          </select>
        </div>
        <div class="mb-3">
          <label for="customer_id">Клиент</label>
          <select v-model="data.customer_id" id="customer_id" name="customer_id" class="form-control">
            <option v-for="customer of customers" :key="customer.id" :value="customer.id">
              {{ customer.full_name }}
            </option>
          </select>
        </div>
        <button class="w-49 btn btn-lg btn-warning">Записать</button>

      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {useRouter} from "vue-router";
import axios from "axios";
import {onMounted, reactive, ref} from "vue";

export default {
  name: "ContractCreate",
  setup() {
    const router = useRouter();
    const data = reactive({
      contract_number: '',
      start_trip_date: '',
      end_trip_date: '',
      agreement_id: '',
      organization_id: '',
      agent_id: '',
      customer_id: '',
    });

    const organizations = ref([])
    const agreements = ref([])
    const agents = ref([])
    const customers = ref([])

    onMounted(async () => {
      const contractOrganization = await axios.get('organizations')
      organizations.value = contractOrganization.data.data;
      const contractAgreement = await axios.get('agreements')
      agreements.value = contractAgreement.data.data
      const contractAgent = await axios.get('agents-list')
      agents.value = contractAgent.data
      const contractCustomer = await axios.get('customers-list')
      customers.value = contractCustomer.data
    })

    const submit = async () => {
      await axios.post('contracts', data)
      await router.push('/contracts')
    }

    return {
      data,
      submit,
      organizations,
      agreements,
      agents,
      customers
    }
  }
}
</script>
