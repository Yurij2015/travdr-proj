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
import {useRouter} from "vue-router";

export default {
  name: "RoleCreate",
  setup() {
    const {push} = useRouter()
    const formData = reactive({
      name: '',
      permissions: [] as string[]
    });

    const permissionList = ref([]);

    onMounted(async () => {
      const {data} = await axios.get('permissions');
      permissionList.value = data;
    });

    const select = (id: number, checked: boolean) => {
      if (checked) {
        formData.permissions = [...formData.permissions, String(id)];
        return;
      }
      formData.permissions = formData.permissions.filter(p => p !== String(id))
    }

    const submit = async () => {
      await axios.post('roles', formData);
      await push('/roles');
    }

    return {
      formData,
      select,
      submit,
      permissionList
    }
  }
}
</script>

<style scoped>

</style>