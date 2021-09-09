<template>
  <form @submit.prevent="submit">
    <h1 class="h3 mb-3 fw-normal">Добавить роль</h1>
    <div class="mb-3 mt-3 row">
      <label class="col-sm-3 col-form-label">Наименование роли:</label>
      <div class="col-sm-9">
        <input name="name" v-model="formData.name" class="form-control" placeholder="Наименование роли">
      </div>
    </div>

    <div class="mb-3 row">
      <label class="col-sm-3 col-form-label">Разрешения:</label>
      <div class="col-sm-9">
        <div class="form-check form-check-inline col-3" v-for="permission in permissionList" :key="permission.id">
          <input type="checkbox" :value="permission.id" class="form-check-input"
                 :checked="checked(permission.id)"
                 @change="select(permission.id, $event.target.checked)">
          <label class="form-check-label">{{ permission.name }}</label>
        </div>
      </div>
    </div>
    <button class="w-100 btn btn-lg btn-primary">Сохранить</button>
  </form>
</template>

<script lang="ts">
import {onMounted, reactive, ref} from "vue"
import axios from "axios";
import {useRoute, useRouter} from "vue-router";
import {Permission} from "@/models/permission";

export default {
  name: "RoleCreate",
  setup() {
    const {push} = useRouter()
    const {params} = useRoute()
    const formData = reactive({
      name: '',
      permissions: [] as string[]
    });

    const permissionList = ref([]);

    onMounted(async () => {
      const {data} = await axios.get('permissions');
      permissionList.value = data;

      const response = await axios.get(`roles/${params.id}`);
      formData.name = response.data.name
      formData.permissions = response.data.permissions.map((p: Permission) => String(p.id))
    });

    const select = (id: number, checked: boolean) => {
      if (checked) {
        formData.permissions = [...formData.permissions, String(id)];
        return;
      }
      formData.permissions = formData.permissions.filter(p => p !== String(id))
    }

    const submit = async () => {
      await axios.put(`roles/${params.id}`, formData);
      await push('/roles');
    }

    // const checked = (id: number) => {
    //   return formData.permissions.some(p => Number(p) === id)
    // }
    const checked = (id: number) => formData.permissions.some(p => Number(p) === id)

    return {
      formData,
      select,
      submit,
      permissionList,
      checked
    }
  }
}
</script>

<style scoped>

</style>