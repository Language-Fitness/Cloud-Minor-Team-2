<script>
import {exercises, headers, generatedQuestions} from "@/views/exercise/exercise";
import axios from "axios";
import {useAuthStore} from "@/stores/store";

const FakeAPI = {
  async fetch({page, itemsPerPage}) {
    return new Promise(resolve => {
      setTimeout(() => {
        const start = (page - 1) * itemsPerPage
        const end = start + itemsPerPage
        const items = exercises.slice()

        const paginated = items.slice(start, end)
        resolve({items: paginated, total: items.length})
      }, 500)
    })
  },
}

export default {
  data: () => ({
    // Show extra functionality if admin
    isAdmin: true,

    //Show generated functionality
    canGenerate: false, // (only if user has access to this)
    questionsToGenerate: 2,
    hasGenerated: false,
    genQuestions: [],

    // Combobox items
    question_types: ['MC'],
    difficulties: ['A1', 'A2', 'B1', 'B2', 'C1', 'C2'],

    // Search filters
    name_search: '',
    name_type_search: '',
    difficulty_search: '',
    soft_deleted_search: false,
    made_by_search: '',

    // Settings for data table
    headers: headers,
    serverItems: [],
    loading: true,
    totalItems: 0,
    itemsPerPage: 10,

    // Open/close dialogs
    dialog: false,
    dialogDelete: false,

    // Item for editing item
    editedIndex: -1,
    editedItem: {
      class_id: '',
      name: '',
      question: '',
      answers: [],
      pos_correct_answer: 0,
      question_type: 'MC',
      difficulty: '',
    },

    // Item for opening 'New Item' tab
    defaultItem: {
      class_id: '',
      name: '',
      question: '',
      answers: [],
      pos_correct_answer: 0,
      question_type: 'MC',
      difficulty: '',
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

      const graphqlEndpoint = 'https://gandalf-the-gateway-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/';

      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const graphqlQuery = `
        query ListExercise($filter: ExerciseFilter!, $paginator: Paginator!) {
          ListExercise(filter: $filter, paginator: $paginator) {
            answers {
              correct
              value
            }
            class_Id
            difficulty
            id
            made_by
            module_id
            name
            question
          }
        }
      `;

      const variables = {
        filter: {
          class_Id: this.$route.params.id
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

        if (data) {
          this.serverItems = data.ListExercise;
          this.totalItems = 1000;
        }
      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }
    },

    editItem(item) {
      this.editedIndex = exercises.findIndex((cl) => cl.id === item.id)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem(item) {
      this.editedIndex = exercises.findIndex((cl) => cl.id === item.id)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
        this.genQuestions = []
        this.hasGenerated = false
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    generate() {
      this.genQuestions = generatedQuestions
      this.hasGenerated = true
      this.editedItem = this.genQuestions[0]
    },

    deleteItemConfirm() {
      // TODO: DELETE LOGIC
      this.closeDelete()
    },

    save() {
      // TODO: SAVE LOGIC
      this.close()
    },

    SaveAndNext() {
      this.editedItem = this.genQuestions[1]
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

        if (data) {
          this.serverItems = data.ListExercise;
          this.totalItems = 1000;
        }
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

      params.class_Id = this.$route.params.id;
      return params
    }
  },
}
</script>

<template>
  <div class="container">
    <header>
      <h1>Exercise overview</h1>
    </header>

    <div class="wrapper">
      <p class="tab-description">
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Delectus hic libero natus nemo obcaecati velit,
        voluptatibus. Aliquam aliquid atque cum cumque fugit hic id, ipsam nulla quisquam sequi. Aut commodi
        consequatur dicta dolorem eos, eveniet fuga magni, nobis officiis possimus quo repellendus, tempora totam!
        Doloremque exercitationem neque repudiandae! Facilis, quae.
      </p>

      <h3 class="my-2 pa-2 w-100" :style="'border: 1px solid lightgray'">Class:
        {{this.$route.params.id}}
      </h3>

      <div class="filter-wrapper w-100 d-flex flex-row">
        <div class="filter-container">
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
                :items="['', 'equals', 'not equals', 'starts', 'ends']"
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
                class="ma-2"
            ></v-combobox>
          </div>

          <!-- ADMIN FILTERS -->
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
            max-width="600px"
            @click:outside="close"
        >
          <!-- New Item Button -->
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

          <v-card>
            <v-card-title class="ml-5 mt-2">
              <span class="text-h5">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text class="ma-0 pt-0">
              <v-container dense>

                <v-row no-gutters v-if="!hasGenerated && this.editedIndex === -1">
                  <v-col
                      cols="6"
                  >
                    <v-combobox
                        class="mb-5"
                        density="compact"
                        v-model="questionsToGenerate"
                        hide-details
                        :items="[2,3,4,5,6,7,8,9,10]"
                        label="How many questions to generate?"
                    ></v-combobox>
                  </v-col>
                  <v-col cols="6" class="pl-2">
                    <v-btn
                        color="blue-darken-1"
                        variant="outlined"
                        @click="generate"
                    >
                      Generate
                    </v-btn>
                  </v-col>
                </v-row>

                <v-divider class="mb-5"></v-divider>

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
                        v-model="editedItem.question"
                        label="Question"
                    ></v-text-field>
                  </v-col>
                </v-row>

                <v-row no-gutters>
                  <v-row>
                    <v-col
                        cols="6"
                    >
                      <v-combobox
                          class="mb-5"
                          v-model="editedItem.question_type"
                          hide-details
                          :items="question_types"
                          label="Question type"
                      ></v-combobox>
                    </v-col>
                    <v-col
                        cols="6"
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
                </v-row>

                <v-row>
                  <v-col cols="6">
                    <v-text-field
                        v-model="todo"
                        label="Answer 1"
                    ></v-text-field>
                    <v-text-field
                        v-model="todo"
                        label="Answer 3"
                    ></v-text-field>
                    <v-text-field
                        v-model="todo"
                        label="Answer 5"
                    ></v-text-field>
                  </v-col>

                  <v-col cols="6">
                    <v-text-field
                        v-model="todo"
                        label="Answer 2"
                    ></v-text-field>
                    <v-text-field
                        v-model="todo"
                        label="Answer 4"
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
                  v-if="hasGenerated"
                  color="blue-darken-1"
                  variant="text"
                  @click="SaveAndNext"
              >
                Save & next
              </v-btn>

              <v-btn
                  v-else
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
