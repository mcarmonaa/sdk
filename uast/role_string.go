// Code generated by "stringer -type=Role"; DO NOT EDIT.

package uast

import "fmt"

const _Role_name = "InvalidSimpleIdentifierQualifiedIdentifierBinaryExpressionBinaryExpressionLeftBinaryExpressionRightBinaryExpressionOpInfixPostfixOpBitwiseLeftShiftOpBitwiseRightShiftOpBitwiseUnsignedRightShiftOpBitwiseOrOpBitwiseXorOpBitwiseAndExpressionStatementOpEqualOpNotEqualOpLessThanOpLessThanEqualOpGreaterThanOpGreaterThanEqualOpSameOpNotSameOpContainsOpNotContainsOpPreIncrementOpPostIncrementOpPreDecrementOpPostDecrementOpNegativeOpPositiveOpBitwiseComplementOpDereferenceOpTakeAddressFileOpBooleanAndOpBooleanOrOpBooleanNotOpBooleanXorOpAddOpSubstractOpMultiplyOpDivideOpModPackageDeclarationImportDeclarationImportPathImportAliasFunctionDeclarationFunctionDeclarationBodyFunctionDeclarationNameFunctionDeclarationReceiverFunctionDeclarationArgumentFunctionDeclarationArgumentNameFunctionDeclarationArgumentDefaultValueFunctionDeclarationVarArgsListTypeDeclarationTypeDeclarationBodyTypeDeclarationBasesTypeDeclarationImplementsVisibleFromInstanceVisibleFromTypeVisibleFromSubtypeVisibleFromPackageVisibleFromSubpackageVisibleFromModuleVisibleFromFriendVisibleFromWorldIfIfConditionIfBodyIfElseSwitchSwitchCaseSwitchCaseConditionSwitchCaseBodySwitchDefaultForForInitForExpressionForUpdateForBodyForEachWhileWhileConditionWhileBodyDoWhileDoWhileConditionDoWhileBodyBreakContinueGotoBlockBlockScopeReturnTryTryBodyTryCatchTryFinallyThrowAssertCallCallReceiverCallCalleeCallPositionalArgumentCallNamedArgumentCallNamedArgumentNameCallNamedArgumentValueNoopBooleanLiteralByteLiteralByteStringLiteralCharacterLiteralListLiteralMapLiteralNullLiteralNumberLiteralRegexpLiteralSetLiteralStringLiteralTupleLiteralTypeLiteralOtherLiteralMapEntryMapKeyMapValueTypePrimitiveTypeAssignmentAssignmentVariableAssignmentValueAugmentedAssignmentAugmentedAssignmentOperatorAugmentedAssignmentVariableAugmentedAssignmentValueThisCommentDocumentationWhitespace"

var _Role_index = [...]uint16{0, 7, 23, 42, 58, 78, 99, 117, 122, 129, 147, 166, 193, 204, 216, 228, 238, 247, 254, 264, 274, 289, 302, 320, 326, 335, 345, 358, 372, 387, 401, 416, 426, 436, 455, 468, 481, 485, 497, 508, 520, 532, 537, 548, 558, 566, 571, 589, 606, 616, 627, 646, 669, 692, 719, 746, 777, 816, 846, 861, 880, 900, 925, 944, 959, 977, 995, 1016, 1033, 1050, 1066, 1068, 1079, 1085, 1091, 1097, 1107, 1126, 1140, 1153, 1156, 1163, 1176, 1185, 1192, 1199, 1204, 1218, 1227, 1234, 1250, 1261, 1266, 1274, 1278, 1283, 1293, 1299, 1302, 1309, 1317, 1327, 1332, 1338, 1342, 1354, 1364, 1386, 1403, 1424, 1446, 1450, 1464, 1475, 1492, 1508, 1519, 1529, 1540, 1553, 1566, 1576, 1589, 1601, 1612, 1624, 1632, 1638, 1646, 1650, 1663, 1673, 1691, 1706, 1725, 1752, 1779, 1803, 1807, 1814, 1827, 1837}

func (i Role) String() string {
	if i < 0 || i >= Role(len(_Role_index)-1) {
		return fmt.Sprintf("Role(%d)", i)
	}
	return _Role_name[_Role_index[i]:_Role_index[i+1]]
}
