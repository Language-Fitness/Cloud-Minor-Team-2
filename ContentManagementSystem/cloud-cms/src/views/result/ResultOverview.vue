<script>
import {headers} from "@/views/result/results";
import axios from "axios";
import {useAuthStore} from "@/stores/store";

export default {
  data: () => ({
    isAdmin: true,
    itemsPerPage: 10,
    headers: headers,
    classId_search: '',
    exerciseId_search: '',

    serverItems: [],
    loading: true,
    totalItems: 0,

    dialog: false,
    dialogDelete: false,
    editedIndex: -1,
  }),
  methods: {
    async loadItems({page, itemsPerPage}) {
      this.loading = true;
      let store = useAuthStore()

      const graphqlEndpoint = 'https://gandalf-the-gateway-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/';

      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const graphqlQuery = `
        query Query($filter: ResultFilter!, $paginator: Paginator!) {
          ListResults(filter: $filter, paginator: $paginator) {
            id
            module_id
            class_id
            exercise_id
            user_id
            input
            result
          }
        }
      `;

      const variables = {
        filter: {
        },
        paginator: {
          Step: 0,
          amount: 10
        }
      }

      try {
        const response = await axios.post(
            graphqlEndpoint,
            {
              query: graphqlQuery,
              variables,
            },
            {headers}
        );

        const {data} = response.data;

        if (data.ListResults) {
          this.serverItems = data.ListResults;
          this.totalItems = 1000;
        }
      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    async save() {
      this.close()
    },

    async filter() {
      let store = useAuthStore()

      const graphqlEndpoint = 'https://gandalf-the-gateway-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/';

      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const graphqlQuery = `
        query Query($filter: ResultFilter!, $paginator: Paginator!) {
          ListResults(filter: $filter, paginator: $paginator) {
            id
            module_id
            class_id
            exercise_id
            user_id
            input
            result
          }
        }
      `;

      let filter = this.buildFilter()
      const variables = {
        filter: filter,
        paginator: {
          Step: 0,
          amount: 10
        }
      }

      try {
        const response = await axios.post(
            graphqlEndpoint,
            {
              query: graphqlQuery,
              variables,
            },
            {headers}
        );

        console.log(response)
        const {data} = response.data;

        if (data.ListResults) {
          this.serverItems = data.ListResults;
        } else {
          this.serverItems = []
        }
        this.totalItems = 1000;

      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }
    },

    buildFilter() {
      const params = {}

      if (this.classId_search !== '') {
        params.classId = this.classId_search;
      }

      if (this.exerciseId_search !== '') {
        params.exerciseId = this.exerciseId_search;
      }
      return params
    }
  }
  ,
}
</script>

<template>
  <div class="container">
    <header>
      <h1>Module overview</h1>
    </header>

    <div class="wrapper">
      <p class="tab-description">
        Op de pagina voor results biedt CRUD-functionaliteit een gestructureerde en efficiënte manier om gegevens
        te beheren. CRUD staat voor Create, Read, Update en Delete, vier fundamentele bewerkingen die essentieel zijn
        voor gegevensbeheer in een systeem. Deze functionaliteit zijn te doen door zowel leraren als admins.
      </p>

      <div class="filter-wrapper d-flex w-100 flex-row">
        <div class="filter-container w-50">
          <h3 class="ml-2">Filter options</h3>

          <div class="d-flex flex-row flex-nowrap w-100">
            <div class="w-50">
            <v-text-field
                :style="'flex-basis: 70%'"
                v-model="exerciseId_search"
                hide-details
                placeholder="Exercise ID..."
                class="mt-2 ml-2"
                density="compact">
            </v-text-field>
            </div>
            <div class="w-50">
            <v-text-field
                :style="'flex-basis: 70%'"
                v-model="classId_search"
                hide-details
                placeholder="Class ID..."
                class="mt-2 ml-2"
                density="compact">
            </v-text-field>
            </div>
          </div>

          <div class="w-50">
            <v-btn @click="filter" type="button" color="primary" class="ma-2">Filter</v-btn>
          </div>
        </div>
      </div>

      <v-data-table-server
          height="100%"
          class="table-entity"
          density="compact"
          v-model:items-per-page="itemsPerPage"
          :headers="headers"
          :items-length="totalItems"
          :items="serverItems"
          :loading="loading"
          item-value="id"
          @update:options="loadItems"
      >
      </v-data-table-server>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 85%;
}

header {
  display: flex;
  flex-direction: column;
  justify-content: center;

  padding: 0 0 0 10px;

  height: 65px;
  color: white;
  background-color: #2a73c5;
}

.wrapper {
  padding: 0 10px 0 10px;
  border: 1px solid lightgray;
}

.tab-description {
  margin: 10px 0 10px 0;
}

.filter-container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-items: center;
  border: 1px solid lightgray;
}

.table-entity {
  margin: 20px 0 0 0;
  border: 1px solid lightgray;
}

.filter-wrapper {
  justify-content: space-between;
  align-items: flex-end;
}
</style>
