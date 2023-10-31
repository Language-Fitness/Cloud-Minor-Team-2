export default class ExcelData {
    // Class definition
    constructor(type, questionText, openQuestionAnswer, multipleChoiceA, multipleChoiceB, multipleChoiceC, multipleChoiceD, multipleChoiceAnswer) {
        this.type = type;
        this.questionText = questionText;
        this.openQuestionAnswer = openQuestionAnswer;
        this.multipleChoiceA = multipleChoiceA;
        this.multipleChoiceB = multipleChoiceB;
        this.multipleChoiceC = multipleChoiceC;
        this.multipleChoiceD = multipleChoiceD;
        this.multipleChoiceAnswer = multipleChoiceAnswer;
    }
}
