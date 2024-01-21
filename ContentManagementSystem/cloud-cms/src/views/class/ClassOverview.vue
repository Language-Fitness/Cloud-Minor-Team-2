<script>
import {headers} from "@/views/class/class";
import {classes} from "@/views/class/class";
import router from '@/router'
import axios from "axios";
import {useAuthStore} from "@/stores/store";

export default {
  data: () => ({
    isAdmin: true,
    itemsPerPage: 10,
    headers: headers,
    difficulties: ['A1', 'A2', 'B1', 'B2', 'C1', 'C2'],
    name_search: '',
    name_type_search: 'eq',
    difficulty_search: '',
    module_id_search: '',
    soft_deleted_search: false,
    made_by_search: '',

    serverItems: [],
    loading: true,
    totalItems: 0,

    dialog: false,
    dialogDelete: false,
    editedIndex: -1,

    editedItem: {
      module_id: '',
      name: '',
      description: '',
      difficulty: '',
    },

    defaultItem: {
      module_id: '',
      name: '',
      description: '',
      difficulty: '',
    },
  }),
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    },
  },
  created() {
    const moduleIdFromQuery = this.$route.query.module;
    if (moduleIdFromQuery) {
      this.defaultItem.module_id = moduleIdFromQuery;
      this.editedItem.module_id = moduleIdFromQuery;
      this.module_id_search = moduleIdFromQuery
    }
  },
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
        query ListClasses($filter: ListClassFilter, $paginate: Paginator) {
          listClasses(filter: $filter, paginate: $paginate) {
            id
            name
            description
            difficulty
            module_Id
            made_by
          }
        }
      `;

      let filter = {}
      if (this.$route.query.module) {
        filter = {
          module_id: this.$route.query.module
        }
      }

      const variables = {
        filter: filter,
        paginate: {
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

        if (data.listClasses) {
          this.serverItems = data.listClasses;
          this.totalItems = 1000;
        }
      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }
    },

    editItem(item) {
      this.editedIndex = classes.findIndex((cl) => cl.id === item.id)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem(item) {
      this.editedIndex = classes.findIndex((cl) => cl.id === item.id)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
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

    deleteItemConfirm() {
      // TODO: DELETE LOGIC
      this.closeDelete()
    },

    async save() {
      let store = useAuthStore()

      const graphqlEndpoint = 'https://gandalf-the-gateway-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/';

      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const graphqlQuery = `
        mutation CreateClass($input: ClassInput!) {
          createClass(input: $input) {
            id
            module_Id
            name
            description
            difficulty
            made_by
            created_at
            updated_at
            soft_deleted
          }
        }
      `;

      const variables = {
        input: {
          name: this.editedItem.name,
          module_Id: this.editedItem.module_id,
          description: this.editedItem.description,
          difficulty: this.editedItem.difficulty,
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
        console.log(data)

      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }

      this.close()
      await this.loadItems({page: 0, itemsPerPage: 10})
    },

    goToClasses(item) {
      router.push('/class/' + item.id + '/exercises?module=' + item.module_Id);
    },

    async filter() {
      let store = useAuthStore()

      const graphqlEndpoint = 'https://gandalf-the-gateway-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/';

      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const graphqlQuery = `
         query ListClasses($filter: ListClassFilter, $paginate: Paginator) {
          listClasses(filter: $filter, paginate: $paginate) {
            id
            name
            description
            difficulty
            module_Id
            made_by
          }
        }
      `;

      let filter = this.buildFilter()
      const variables = {
        filter: filter,
        paginate: {
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

        if (data.listClasses) {
          this.serverItems = data.listClasses;
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

      if (this.name_search !== '') {
        params.name = {
          type: this.name_type_search,
          input: this.name_search
        }
      }

      if (this.module_id_search !== '') {
        params.module_Id = this.module_id_search;
      }

      if (this.difficulty_search !== '') {
        params.difficulty = this.difficulty_search;
      }
      return params
    }
  },
}
</script>

<template>
  <div class="container">
    <header>
      <h1>Class overview</h1>
    </header>

    <div class="wrapper">
      <p class="tab-description">
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Delectus hic libero natus nemo obcaecati velit,
        voluptatibus. Aliquam aliquid atque cum cumque fugit hic id, ipsam nulla quisquam sequi. Aut commodi
        consequatur dicta dolorem eos, eveniet fuga magni, nobis officiis possimus quo repellendus, tempora totam!
        Doloremque exercitationem neque repudiandae! Facilis, quae.
      </p>

      <div class="filter-wrapper w-100 d-flex flex-row">
        <div class="filter-container w-50">
          <h3 class="ml-2">Filter options</h3>

          <div class="d-flex flex-row flex-nowrap w-100">
            <v-text-field
                :style="'flex-basis: 70%'"
                v-model="name_search"
                hide-details
                placeholder="Name..."
                class="mt-2 ml-2"
                density="compact">
            </v-text-field>
            <v-combobox
                v-model="name_type_search"
                hide-details
                :items="['eq', 'ne', 'starts', 'ends']"
                density="compact"
                class="mt-2 mr-2 ml-1"
            ></v-combobox>
          </div>

          <div class="w-50">
            <v-combobox
                v-model="difficulty_search"
                hide-details
                :items="difficulties"
                density="compact"
                label="Difficulty"
                class="mx-2 mt-2"
            ></v-combobox>
          </div>

          <!-- ADMIN FILTERS -->
          <div class="w-50">
            <v-text-field
                v-if="isAdmin"
                v-model="module_id_search"
                hide-details
                placeholder="Module ID..."
                class="mx-2 mt-2"
                density="compact">
            </v-text-field>
          </div>

          <div class="w-50">
            <v-combobox
                v-if="isAdmin"
                v-model="soft_deleted_search"
                disabled
                hide-details
                :items="[true, false]"
                density="compact"
                label="Soft deleted"
                class="ma-2"
            ></v-combobox>
          </div>

          <div class="w-50">
            <v-text-field
                v-if="isAdmin"
                v-model="made_by_search"
                disabled
                hide-details
                placeholder="Made by..."
                class="ma-2"
                density="compact">
            </v-text-field>
          </div>

          <div class="w-50">
            <v-btn @click="filter" type="button" color="primary" class="ma-2">Filter</v-btn>
          </div>
        </div>
        <!-- MODALS -->
        <v-dialog
            v-model="dialog"
            max-width="400px"
        >
          <template v-slot:activator="{ props }">
            <v-btn
                color="primary"
                class="ml-2"
                dark
                v-bind="props"
                border
            >
              New Item
            </v-btn>

          </template>

          <!-- EDITING SECTION -->
          <v-card>
            <v-card-title class="ml-5 mt-2">
              <span class="text-h5">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text class="ma-0 pt-0">
              <v-container dense>
                <v-row no-gutters>
                  <v-col
                      cols="12"
                  >
                    <v-text-field
                        v-model="editedItem.name"
                        label="Name"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row no-gutters>
                  <v-col
                      cols="12"
                  >
                    <v-text-field
                        v-model="editedItem.module_id"
                        label="module id"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row no-gutters>
                  <v-col
                      cols="12"
                  >
                    <v-text-field
                        v-model="editedItem.description"
                        label="description"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row no-gutters>
                  <v-col
                      cols="12"
                  >
                    <v-combobox
                        class="mb-5"
                        v-model="editedItem.difficulty"
                        hide-details
                        :items="difficulties"
                        label="Difficulty"
                    ></v-combobox>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <!-- EDITING CONFIRMATION SECTION -->
            <v-card-actions class="flex-row justify-space-between mx-5">
              <v-btn
                  color="blue-darken-1"
                  variant="text"
                  @click="close"
              >
                Cancel
              </v-btn>
              <v-btn
                  color="blue-darken-1"
                  variant="text"
                  @click="save"
              >
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>

      <!-- DELETION MODAL -->
      <v-dialog v-model="dialogDelete" max-width="500px">
        <v-card class="flex-column align-center pa-5">
          <v-card-title class="text-h5">Are you sure you want to delete this item?</v-card-title>

          <v-checkbox v-if="isAdmin" class="pa-0" hide-details label="Hard delete"></v-checkbox>

          <v-card-text class="pt-0">You will not be able to recover this item.</v-card-text>

          <v-card-actions class="flex-row justify-space-between">
            <v-btn color="blue-darken-1" variant="text" @click="closeDelete">Cancel</v-btn>
            <v-btn color="blue-darken-1" variant="text" @click="deleteItemConfirm">OK</v-btn>
            <v-spacer></v-spacer>
          </v-card-actions>
        </v-card>
      </v-dialog>

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
        <template v-slot:item.actions="{ item }">
          <v-icon
              size="small"
              class="me-2"
              @click="editItem(item)"
          >
            mdi-pencil
          </v-icon>
          <v-icon
              size="small"
              @click="deleteItem(item)"
          >
            mdi-delete
          </v-icon>
        </template>

        <template v-slot:item.exercises="{ item }">
          <v-icon
              size="small"
              class="me-2"
              @click="goToClasses(item)"
          >
            mdi-arrow-right-bold-circle-outline
          </v-icon>
        </template>
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
