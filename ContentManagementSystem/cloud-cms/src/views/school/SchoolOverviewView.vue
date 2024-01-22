<script>
import {headers, listSchoolsQuery} from "@/views/school/school";
import {useAuthStore} from "@/stores/store";
import axios from "axios";

export default {
  data: () => ({
    isAdmin: true,
    itemsPerPage: 10,
    headers: headers,
    name_search: '',
    name_type_search: '',
    location_search: '',
    soft_deleted_search: false,
    made_by_search: '',

    serverItems: [],
    loading: true,
    totalItems: 0,

    dialog: false,
    dialogDelete: false,
    editedIndex: -1,

    editedItem: {
      name: '',
      location: '',
    },

    defaultItem: {
      name: '',
      location: '',
    },
  }),
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    },
  },
  methods: {
    async loadItems({page, itemsPerPage}) {
      this.loading = true;
      let store = useAuthStore()

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;;
      const graphqlQuery = listSchoolsQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const variables = {
        filter: {
        },
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

        if (data.listSchools) {
          this.serverItems = data.listSchools;
          this.totalItems = 1000;
        }
      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }
    },

    editItem(item) {
      this.editedIndex = schools.findIndex((school) => school.id === item.id)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem(item) {
      this.editedIndex = schools.findIndex((school) => school.id === item.id)
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
      this.closeDelete()
    },

    save() {
      this.close()
    },
  },
}
</script>

<template>
  <div class="container">
    <header>
      <h1>School overview</h1>
    </header>

    <div v-if="isAdmin" class="wrapper">
      <p class="tab-description">
        Op de pagina voor scholen biedt CRUD-functionaliteit een gestructureerde en efficiënte manier om gegevens
        te beheren. CRUD staat voor Create, Read, Update en Delete, vier fundamentele bewerkingen die essentieel zijn
        voor gegevensbeheer in een systeem. Deze functionaliteit zijn te doen door admins.
      </p>

      <div class="filter-wrapper d-flex w-100 flex-row">
      <div class="filter-container w-500">
        <h2 class="ml-2">Filter options</h2>
        <v-divider></v-divider>

        <div class="d-flex flex-row flex-nowrap w-100">
          <v-text-field
              v-if="isAdmin"
              :style="'flex-basis: 70%'"
              v-model="name_search"
              hide-details
              placeholder="Name..."
              class="mt-2 ml-2"
              density="compact">
          </v-text-field>
          <v-combobox
              v-if="isAdmin"
              v-model="name_type_search"
              hide-details
              :items="['', 'equals', 'not equals', 'starts', 'ends']"
              density="compact"
              class="mt-2 mr-2 ml-1"
          ></v-combobox>
        </div>

        <div class="w-50">
          <v-text-field
              v-if="isAdmin"
              v-model="location_search"
              hide-details
              placeholder="Location..."
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
              class="mx-2 mt-2"
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
          <v-btn type="button" color="primary" class="ma-2">Filter</v-btn>
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
                        v-model="editedItem.location"
                        label="location"
                    ></v-text-field>
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