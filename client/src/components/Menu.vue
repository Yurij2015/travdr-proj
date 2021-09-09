<template>
  <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse mt-3">
    <div class="position-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link to="/" class="nav-link" active-class="active">
            Панель управления
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/users" class="nav-link" active-class="active">
            Пользователи
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/roles" class="nav-link" active-class="active">
            Роли
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/customers" class="nav-link" active-class="active">
            Клиенты
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/agents" class="nav-link" active-class="active">
            Агенты
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/contracts" class="nav-link" active-class="active">
            Договора
          </router-link>
        </li>        <li class="nav-item">
        <router-link to="/agreements" class="nav-link" active-class="active">
          Соглашения
        </router-link>
      </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: "Menu"
}
</script>

<style scoped>

</style>