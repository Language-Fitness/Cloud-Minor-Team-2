<script>
import {
  createExerciseQuery,
  deleteExerciseQuery, difficulties,
  generateMcQuery,
  headers,
  listExercisesQuery, retrieveMcQuery
} from "@/views/exercise/exercise";
import axios from "axios";
import {useAuthStore} from "@/stores/store";

export default {
  data: () => ({
    // Show extra functionality if admin
    isAdmin: true,
    categories: ['grammatica', 'spelling', 'woordenschat', 'uitdrukkingen', 'interpunctie', 'werkwoordvervoegingen'],
    //Show generated functionality
    isGenerating: false,

    questionsToGenerate: 2,
    questionDifficulty: 'B2',
    questionSubject: 'woordenschat',

    hasGenerated: false,
    genQuestions: [],

    // Combobox items
    difficulties: difficulties,

    // Search filters
    name_search: '',
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
    exerciseToDelete: '',

    // Item for editing item
    editedIndex: -1,
    editedItem: {
      class_Id: '',
      module_Id: '',
      name: '',
      question: '',
      answers: ['', '', '', '', ''],
      question_type: 'MC',
      difficulty: '',
    },

    // Item for opening 'New Item' tab
    defaultItem: {
      class_Id: '',
      module_Id: '',
      name: '',
      question: '',
      answers: ['', '', '', '', ''],
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

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;
      const graphqlQuery = listExercisesQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

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

        if (data.ListExercise) {
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
      this.dialog = true
    },

    deleteItem(item) {
      this.exerciseToDelete = item.id
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

    async generate() {
      let token = await this.generateMc()
      this.isGenerating = true

      const intervalId = setInterval(async () => {
        let res = await this.retrieveMc(token);

        if (res.message !== 'Response still pending, please wait.') {
          clearInterval(intervalId)

          this.hasGenerated = true
          this.genQuestions = res.questions
          this.isGenerating = false
          this.setEditForm()
        }
      }, 5000);
    },

    setEditForm() {
      this.editedItem.answers = []
      this.editedItem.name = this.genQuestions[0].questionSubject
      this.editedItem.question = this.genQuestions[0].questionText
      this.editedItem.difficulty = this.genQuestions[0].questionLevel
      this.editedItem.class_Id = this.$route.params.id
      this.editedItem.module_Id = this.$route.query.module;

      this.genQuestions[0].answerOptions.forEach((element, index) => {
        this.editedItem.answers.push({
          correct: element === this.genQuestions[0].correctAnswer,
          value: element
        })
      });
    },

    async deleteItemConfirm() {
      let store = useAuthStore()

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;
      const graphqlQuery = deleteExerciseQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const variables = {
        filter: {
          object_id: this.exerciseToDelete,
          object_type: "Exercise"
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

      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }

      await this.loadItems({page: 0, itemsPerPage: 10})
      this.closeDelete()
    },

    async save() {
      let store = useAuthStore()

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;
      const graphqlQuery = createExerciseQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const variables = {
        exercise: {
          class_Id: this.editedItem.class_Id,
          module_id: this.editedItem.module_Id,
          name: this.editedItem.name,
          difficulty: this.editedItem.difficulty,
          question: this.editedItem.question,
          answers: this.editedItem.answers
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
        this.genQuestions.shift()

      } catch (error) {
        console.error('GraphQL request failed', error);
      } finally {
        this.loading = false;
      }
    },

    async SaveAndNext() {
      if (this.genQuestions.length === 1) {
        await this.save()
        await this.loadItems({page: 0, itemsPerPage: 10})
        this.close()
      } else {
        await this.save()
        this.setEditForm()
      }
    },

    async filter() {
      let store = useAuthStore()

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;
      const graphqlQuery = listExercisesQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

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

        const {data} = response.data;
        console.log(response)

        if (data.ListExercise) {
          this.serverItems = data.ListExercise;
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
        params.name = this.name_search
      }

      if (this.difficulty_search !== '') {
        params.difficulty = this.difficulty_search;
      }

      params.class_Id = this.$route.params.id;
      return params
    },

    async generateMc() {
      let store = useAuthStore()

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;
      const graphqlQuery = generateMcQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      try {
        const response = await axios.post(
            graphqlEndpoint,
            {
              query: graphqlQuery,
              variables: {
                amountQuestions: this.questionsToGenerate,
                questionLevel: this.questionDifficulty,
                questionSubject: this.questionSubject,
              },
            },
            {headers}
        );
        console.log(response)
        const {data} = response.data;
        return data.generateMultipleChoiceQuestions.response.token

      } catch (error) {
        console.error('GraphQL request failed', error);
      }
    },

    async retrieveMc(token) {
      let store = useAuthStore()

      const graphqlEndpoint = import.meta.env.VITE_GATEWAY_ENDPOINT;
      const graphqlQuery = retrieveMcQuery;
      const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${store.accessToken}`,
      };

      const variables = {
        token: token,
      }

      try {
        const response = await axios.post(
            graphqlEndpoint,
            {
              query: graphqlQuery,
              variables
            },
            {headers}
        );
        console.log(response)
        const {data} = response.data;
        return {message: data.retrieveMultipleChoiceQuestions.message, questions: data.retrieveMultipleChoiceQuestions.questions}

      } catch (error) {
        console.error('GraphQL request failed', error);
      }
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
        Op de pagina voor exercises biedt CRUD-functionaliteit een gestructureerde en efficiënte manier om gegevens
        te beheren. CRUD staat voor Create, Read, Update en Delete, vier fundamentele bewerkingen die essentieel zijn
        voor gegevensbeheer in een systeem. Deze functionaliteit zijn te doen door zowel leraren als admins.
      </p>

      <h3 class="my-2 pa-2" :style="'border: 1px solid lightgray'">Class:
        {{this.$route.params.id}}
      </h3>

      <div class="filter-wrapper w-100 d-flex flex-row">
        <div class="filter-container">
          <h2 class="ml-2">Filter options</h2>
          <v-divider></v-divider>

          <div class="d-flex flex-row flex-nowrap w-100">
            <v-text-field
                :style="'flex-basis: 70%'"
                v-model="name_search"
                hide-details
                placeholder="Name..."
                class="mt-2 ml-2"
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

                <v-row v-if="!hasGenerated && this.editedIndex === -1">
                  <v-col
                      cols="6"
                  >
                    <v-combobox
                        :disabled="isGenerating"
                        class="mb-5"
                        v-model="questionsToGenerate"
                        hide-details
                        :items="[2,3,4,5,6,7,8,9,10]"
                        label="How many questions to generate?"
                    ></v-combobox>
                  </v-col>

                  <v-col
                      cols="6"
                  >
                    <v-combobox
                        :disabled="isGenerating"
                        class="mb-5"
                        v-model="questionDifficulty"
                        hide-details
                        :items="difficulties"
                        label="Difficulty"
                    ></v-combobox>
                  </v-col>
                </v-row>

                <v-row v-if="!hasGenerated && this.editedIndex === -1">
                  <v-col cols="6">
                    <v-combobox
                        :disabled="isGenerating"
                        class="mb-5"
                        v-model="questionSubject"
                        hide-details
                        :items="categories"
                        label="Subject"
                    ></v-combobox>
                  </v-col>


                  <v-col cols="6" class="pl-2">
                    <v-btn
                        :disabled="isGenerating"
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
                        :disabled="isGenerating"
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
                        :disabled="isGenerating"
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
                          :disabled="isGenerating"
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
                  <v-col cols="12">
                    <v-row>
                      <v-col v-for="(answer, index) in editedItem.answers" :key="index" cols="6">
                        <v-text-field :disabled="isGenerating" v-model="answer.value" :label="'Answer ' + (index + 1)"></v-text-field>
                      </v-col>
                    </v-row>
                  </v-col>

                </v-row>
              </v-container>
            </v-card-text>

            <!-- EDITING CONFIRMATION SECTION -->
            <v-card-actions class="flex-row justify-space-between mx-5">
              <v-btn
                  :disabled="isGenerating"
                  color="blue-darken-1"
                  variant="text"
                  @click="close"
              >
                Cancel
              </v-btn>

              <v-btn
                  :disabled="isGenerating"
                  v-if="hasGenerated"
                  color="blue-darken-1"
                  variant="text"
                  @click="SaveAndNext"
              >
                Save & next
              </v-btn>

              <v-btn
                  :disabled="isGenerating"
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
