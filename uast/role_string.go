// Code generated by "stringer -type=Role"; DO NOT EDIT.

package uast

import "fmt"

const _Role_name = "SimpleIdentifierQualifiedIdentifierBinaryExpressionBinaryExpressionLeftBinaryExpressionRightBinaryExpressionOpInfixPostfixOpBitwiseLeftShiftOpBitwiseRightShiftOpBitwiseUnsignedRightShiftOpBitwiseOrOpBitwiseXorOpBitwiseAndExpressionStatementOpEqualOpNotEqualOpLessThanOpLessThanEqualOpGreaterThanOpGreaterThanEqualOpSameOpNotSameOpContainsOpNotContainsOpPreIncrementOpPostIncrementOpPreDecrementOpPostDecrementOpNegativeOpPositiveOpBitwiseComplementOpDereferenceOpTakeAddressFileOpBooleanAndOpBooleanOrOpBooleanNotOpBooleanXorOpAddOpSubstractOpMultiplyOpDivideOpModPackageDeclarationImportDeclarationImportPathImportAliasFunctionDeclarationFunctionDeclarationBodyFunctionDeclarationNameFunctionDeclarationReceiverFunctionDeclarationArgumentFunctionDeclarationArgumentNameFunctionDeclarationArgumentDefaultValueFunctionDeclarationVarArgsListTypeDeclarationTypeDeclarationBodyTypeDeclarationBasesTypeDeclarationImplementsVisibleFromInstanceVisibleFromTypeVisibleFromSubtypeVisibleFromPackageVisibleFromSubpackageVisibleFromModuleVisibleFromFriendVisibleFromWorldIfIfConditionIfBodyIfElseSwitchSwitchCaseSwitchCaseConditionSwitchCaseBodySwitchDefaultForForInitForExpressionForUpdateForBodyForEachWhileWhileConditionWhileBodyDoWhileDoWhileConditionDoWhileBodyBreakContinueGotoBlockBlockScopeReturnTryTryBodyTryCatchTryFinallyThrowAssertCallCallReceiverCallCalleeCallPositionalArgumentCallNamedArgumentCallNamedArgumentNameCallNamedArgumentValueNoopBooleanLiteralByteLiteralByteStringLiteralCharacterLiteralListLiteralMapLiteralNullLiteralNumberLiteralRegexpLiteralSetLiteralStringLiteralTupleLiteralTypeLiteralOtherLiteralMapEntryMapKeyMapValueTypePrimitiveTypeAssignmentAssignmentVariableAssignmentValueAugmentedAssignmentAugmentedAssignmentOperatorAugmentedAssignmentVariableAugmentedAssignmentValueThisCommentDocumentationWhitespace"

var _Role_index = [...]uint16{0, 16, 35, 51, 71, 92, 110, 115, 122, 140, 159, 186, 197, 209, 221, 231, 240, 247, 257, 267, 282, 295, 313, 319, 328, 338, 351, 365, 380, 394, 409, 419, 429, 448, 461, 474, 478, 490, 501, 513, 525, 530, 541, 551, 559, 564, 582, 599, 609, 620, 639, 662, 685, 712, 739, 770, 809, 839, 854, 873, 893, 918, 937, 952, 970, 988, 1009, 1026, 1043, 1059, 1061, 1072, 1078, 1084, 1090, 1100, 1119, 1133, 1146, 1149, 1156, 1169, 1178, 1185, 1192, 1197, 1211, 1220, 1227, 1243, 1254, 1259, 1267, 1271, 1276, 1286, 1292, 1295, 1302, 1310, 1320, 1325, 1331, 1335, 1347, 1357, 1379, 1396, 1417, 1439, 1443, 1457, 1468, 1485, 1501, 1512, 1522, 1533, 1546, 1559, 1569, 1582, 1594, 1605, 1617, 1625, 1631, 1639, 1643, 1656, 1666, 1684, 1699, 1718, 1745, 1772, 1796, 1800, 1807, 1820, 1830}

func (i Role) String() string {
	i -= 1
	if i < 0 || i >= Role(len(_Role_index)-1) {
		return fmt.Sprintf("Role(%d)", i+1)
	}
	return _Role_name[_Role_index[i]:_Role_index[i+1]]
}
