<script>
import {headers, schools} from "@/views/school/school";

const FakeAPI = {
  async fetch({page, itemsPerPage}) {
    return new Promise(resolve => {
      setTimeout(() => {
        const start = (page - 1) * itemsPerPage
        const end = start + itemsPerPage
        const items = schools.slice()

        const paginated = items.slice(start, end)
        resolve({items: paginated, total: items.length})
      }, 500)
    })
  },
}

export default {
  data: () => ({
    isAdmin: true,
    itemsPerPage: 10,
    headers: headers,
    categories: ['Grammatica', 'Spelling', 'Woordenschat', 'Werkwoordspelling', 'Fastlane'],
    difficulties: ['A1', 'A2', 'B1', 'B2', 'C1', 'C2'],
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
    loadItems({page, itemsPerPage}) {
      this.loading = true
      FakeAPI.fetch({page, itemsPerPage}).then(({items, total}) => {
        this.serverItems = items
        this.totalItems = total
        this.loading = false
      })
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
      // TODO: DELETE LOGIC
      this.closeDelete()
    },

    save() {
      // TODO: SAVE LOGIC
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
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Delectus hic libero natus nemo obcaecati velit,
        voluptatibus. Aliquam aliquid atque cum cumque fugit hic id, ipsam nulla quisquam sequi. Aut commodi
        consequatur dicta dolorem eos, eveniet fuga magni, nobis officiis possimus quo repellendus, tempora totam!
        Doloremque exercitationem neque repudiandae! Facilis, quae.
      </p>

      <div class="filter-container">

        <div class="d-flex flex-row flex-nowrap w-50">
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

        <div class="w-25">
          <v-text-field
              v-if="isAdmin"
              v-model="location_search"
              hide-details
              placeholder="Location..."
              class="mx-2 mt-2"
              density="compact">
          </v-text-field>
        </div>

        <div class="w-25">
          <v-combobox
              v-if="isAdmin"
              v-model="soft_deleted_search"
              hide-details
              :items="[true, false]"
              density="compact"
              label="Soft deleted"
              class="ma-2"
          ></v-combobox>
        </div>

        <div class="w-25">
          <v-text-field
              v-if="isAdmin"
              v-model="made_by_search"
              hide-details
              placeholder="Made by..."
              class="ma-2"
              density="compact">
          </v-text-field>
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
</style>