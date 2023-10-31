<template>
  <div class="excel-reader">
    <div class="upload-container">
      <input
          type="file"
          @change="handleFileUpload"
          class="file-input"
          ref="fileInput"
          style="display: none"
      />
      <label class="file-label" @click="openFileInput">
        <span class="file-icon">📂</span> Choose an Excel file
      </label>
    </div>
    <div v-if="excelData">
      <!-- Display the content of the uploaded Excel file here -->
      <table class="excel-table">
        <thead>
        <tr>
          <th v-for="(header, index) in excelData[0]" :key="index">{{ header }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(row, rowIndex) in excelData.slice(1)" :key="rowIndex">
          <td v-for="(cell, cellIndex) in row" :key="cellIndex">{{ cell }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import * as XLSX from 'xlsx';
import ExcelData from '../Models/ExcelData';

export default {
  data() {
    return {
      excelData: null,
      excelDataList: [], // Array to store ExcelData instances
    };
  },
  methods: {
    openFileInput() {
      this.$refs.fileInput.click(); // Trigger the hidden file input element
    },
    handleFileUpload(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();

        reader.onload = (e) => {
          const data = new Uint8Array(e.target.result);
          const workbook = XLSX.read(data, { type: 'array' });
          const sheetName = workbook.SheetNames[0];
          const sheet = workbook.Sheets[sheetName];
          this.excelData = XLSX.utils.sheet_to_json(sheet, { header: 1 });
          this.excelDataList = this.excelData.slice(1).map(row => {
            return new ExcelData(
                row[0],
                row[1],
                row[2],
                row[3],
                row[4],
                row[5],
                row[6],
                row[7]
            );
          });

          // Log the excelData object
          console.log(this.excelData);
        };

        reader.readAsArrayBuffer(file);
      }
    },
  },
};
</script>

<style scoped>
.excel-reader {
  text-align: center;
  padding: 20px;
}

.upload-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.file-input {
  display: none;
}

.file-label {
  display: inline-block;
  background-color: #3498db;
  color: #fff;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.2s;
}

.file-label:hover {
  background-color: #2980b9;
}

.file-icon {
  margin-right: 10px;
}

.excel-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.excel-table th, .excel-table td {
  border: 1px solid #ccc;
  padding: 10px;
}

.excel-table th {
  background-color: #3498db;
  color: #fff;
  font-weight: bold;
}

.excel-table td {
  background-color: #f9f9f9;
}
</style>
