<script>
import {exercises, headers, generatedQuestions} from "@/views/exercise/exercise";
import axios from "axios";

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
    question_search: '',
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

      const graphqlEndpoint = 'https://gandalf-the-gateway-bramterlouw-dev.apps.ocp2-inholland.joran-bergfeld.com/';
      const accessToken = 'eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDQ5OTAzOTQsImlhdCI6MTcwNDk4OTQ5NCwianRpIjoiYjFlM2M3MGUtNDI3Ni00NTk1LTliN2ItOWEzZTM4MGRhYTg2IiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiNmMxY2U0NDgtNjcwZi00N2IyLTgzZjctNGQ3NzFiMDE3NzViIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6ImM0ZTYwMTg2LWQ4MTctNDk3Zi04Yzk0LTVhNTU4NjAwNmNiNSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9yZXN1bHRfc29mdERlbGV0ZSIsImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfc2Nob29sX25hbWUiLCJ1cGRhdGVfcmVzdWx0IiwiZmlsdGVyX2V4ZXJjaXNlX21vZHVsZV9pZCIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwiZmlsdGVyX21vZHVsZV9kaWZmaWN1bHR5IiwiZmlsdGVyX3Jlc3VsdF9tb2R1bGVfaWQiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwiZmlsdGVyX3NjaG9vbF9zb2Z0RGVsZXRlIiwiZGVsZXRlX3Jlc3VsdF9hbGwiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJjcmVhdGVfcmVzdWx0IiwiZ2V0X3Jlc3VsdF9hbGwiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnkiLCJsaXN0X3Jlc3VsdHNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX3F1ZXN0aW9uX3R5cGVfaWQiLCJ1cGRhdGVfY2xhc3NfYWxsIiwiZ2V0X2NsYXNzIiwiZ2V0X3NjaG9vbHNfYWxsIiwiZmlsdGVyX3Jlc3VsdF9leGVyY2lzZV9pZCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwidXBkYXRlX3Jlc3VsdF9hbGwiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zX2Zyb21fZmlsZSIsImdldF9jbGFzc2VzX2FsbCIsInVwZGF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX21hZGVfYnkiLCJnZXRfZXhlcmNpc2VzX2FsbCIsIm9wZW5haV9nZW5lcmF0ZV9leHBsYW5hdGlvbiIsImZpbHRlcl9jbGFzc19tYWRlX2J5IiwiZmlsdGVyX21vZHVsZV9zb2Z0RGVsZXRlIiwiZ2V0X2V4ZXJjaXNlcyIsImdldF9jbGFzc2VzIiwiZGVsZXRlX21vZHVsZSIsImdldF9zY2hvb2xzIiwiZGVsZXRlX2V4ZXJjaXNlIiwidXBkYXRlX2V4ZXJjaXNlIiwiZ2V0X2V4ZXJjaXNlIiwiZmlsdGVyX3Jlc3VsdF91c2VyX2lkIiwiZmlsdGVyX2V4ZXJjaXNlX25hbWUiLCJmaWx0ZXJfZXhlcmNpc2Vfc29mdERlbGV0ZSIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfcmVzdWx0X2NsYXNzX2lkIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJkZWxldGVfcmVzdWx0IiwiY3JlYXRlX21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJjcmVhdGVfY2xhc3MiLCJjcmVhdGVfc2Nob29sIiwiZ2V0X21vZHVsZXNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX2NsYXNzX2lkIiwibGlzdF9yZXN1bHRzIiwiZmlsdGVyX21vZHVsZV9zY2hvb2xfaWQiLCJmaWx0ZXJfY2xhc3NfbmFtZSIsImdldF9yZXN1bHQiLCJmaWx0ZXJfc2Nob29sX2hhc19vcGVuYWlfYWNjZXNzIiwib3BlbmFpX2dldF9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlIiwiZmlsdGVyX21vZHVsZV9uYW1lIiwiZmlsdGVyX21vZHVsZV9tYWRlX2J5X25hbWUiLCJmaWx0ZXJfZXhlcmNpc2VfbWFkZV9ieSIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZmlsdGVyX21vZHVsZV9wcml2YXRlIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiYzRlNjAxODYtZDgxNy00OTdmLThjOTQtNWE1NTg2MDA2Y2I1IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.BMOaN_asRJV0OjqbIMP01Gcp7VMoZn689UOMwMZG3lZABGdm76gKTcQVKLMURJjtr6znnSBDIoomCqUS0HTiIRwkwpIRfYC-q-CvOoDtesNDBzbfC5lvYhDbEycpGU-xoaK11cp5WcGAYFzyGxN-3pZjUAamNNitW3ENgYK9LCpa3gMWLCmF7LUZPuvE8lzjQE0EOxvncjfYR5Kux1u0eJvhf31krOH0sRsbB_-4O1qaPLIuFxKRnHbUP73VOsY60S2f8cveH5nFBahORkWmxBd_LKChMwc1IFgzFP0QkidA6HKw7KXViXuy5LjgzSoP5B1MoIrMpbPICaKdsEVGTw';

      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${accessToken}`,
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
        filter: {},
        paginator: {
          Step: 0,
          amount: 10
        }
      }

      // try {
      //   const response = await axios.post(
      //       graphqlEndpoint,
      //       {
      //         query: graphqlQuery,
      //         variables,
      //       },
      //       {headers}
      //   );
      //   const {data} = response.data;
      //   console.log(data)
      //
      //   this.serverItems = data.ListExercise;
      //   this.totalItems = 1000;
      // } catch (error) {
      //   console.error('GraphQL request failed', error);
      // } finally {
      //   this.loading = false;
      // }
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
        2055b38e-992d-11ee-b9d1-0242ac120002
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
            <v-text-field
                v-model="question_search"
                hide-details
                placeholder="Question..."
                class="mt-2 ml-2 mr-2"
                density="compact">
            </v-text-field>
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
                hide-details
                placeholder="Made by..."
                class="ma-2"
                density="compact">
            </v-text-field>
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
