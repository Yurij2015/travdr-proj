import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
// import Users from "@/pages/users/Users.vue";
import Register from "@/pages/Register.vue";
import Wrapper from "@/pages/Wrapper.vue";
import Login from "@/pages/Login.vue";
import UserCreate from "@/pages/users/UserCreate.vue";
import UserEdit from "@/pages/users/UserEdit.vue";
import Roles from "@/pages/roles/Roles.vue";
import RoleCreate from "@/pages/roles/RoleCreate.vue";
import RoleEdit from "@/pages/roles/RoleEdit.vue";
import CustomerCreate from "@/pages/customers/CustomerCreate.vue";
import CustomerEdit from "@/pages/customers/CustomerEdit.vue";
import Profile from "@/pages/Profile.vue";
import AgentCreate from "@/pages/agents/AgentCreate.vue"
import AgentEdit from "@/pages/agents/AgentEdit.vue"
import AgentCard from "@/pages/agents/AgentCard.vue";
import ContractCreate from "@/pages/contracts/ContractCreate.vue";
import ContractEdit from "@/pages/contracts/ContractEdit.vue";
import AgreementCreate from "@/pages/agreements/AgreementCreate.vue";
import AgreementEdit from "@/pages/agreements/AgreementEdit.vue";
import ContractCard from "@/pages/contracts/ContractCard.vue";
import AgreementCard from "@/pages/agreements/AgreementCard.vue";


const routes: Array<RouteRecordRaw> = [
    {path: "/register", component: Register},
    {path: "/login", component: Login},
    {
        path: '',
        component: Wrapper,
        children: [
            {path: '', component: () => import(("@/pages/Dashboard"))},
            {path: '/users/create', component: UserCreate},
            {path: '/users/:id/edit', component: UserEdit},
            {path: '/roles', component: Roles},
            {path: '/agents', component: () => import(("@/pages/agents/Agents"))},
            {path: '/agents/create', component: AgentCreate},
            {path: '/agents/:id/edit', component: AgentEdit},
            {path: '/agents/:id/card', component: AgentCard},

            {path: '/roles/create', component: RoleCreate},
            {path: '/roles/:id/edit', component: RoleEdit},
            {path: '/customers', component: () => import(("@/pages/customers/Customers"))},

            {path: '/users', component: () => import(("@/pages/users/Users"))},
            {path: '/customers/create', component: CustomerCreate},
            {path: '/customers/:id/edit', component: CustomerEdit},
            {path: '/profile', component: Profile},

            {path: '/contracts', component: () => import(("@/pages/contracts/Contracts"))},
            {path: '/contracts/create', component: ContractCreate},
            {path: '/contracts/:id/edit', component: ContractEdit},
            {path: '/contracts/:id/card', component: ContractCard},


            {path: '/agreements', component: () => import(("@/pages/agreements/Agreements"))},
            {path: '/agreements/create', component: AgreementCreate},
            {path: '/agreements/:id/edit', component: AgreementEdit},
            {path: '/agreements/:id/card', component: AgreementCard},


        ]
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
