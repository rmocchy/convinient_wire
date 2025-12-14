package app

// StructAnalysisResult は構造体の解析結果を保持する
type StructAnalysisResult struct {
	StructName    string             // 構造体名
	PackagePath   string             // パッケージパス
	InitFunctions []InitFunctionInfo // 構造体を返す初期化関数
	Fields        []FieldNode        // フィールドのノード (StructNode または InterfaceNode)
	Skipped       bool               // 解析がスキップされたかどうか
	SkipReason    string             // スキップされた理由
}

// InitFunctionInfo は初期化関数の情報を保持する
type InitFunctionInfo struct {
	Name        string // 関数名
	PackagePath string // パッケージパス
}

// FieldNode はフィールドを表すインターフェース
type FieldNode interface {
	GetFieldName() string
	IsStructNode() bool
	IsInterfaceNode() bool
}

// StructNode は構造体フィールドを表す
type StructNode struct {
	FieldName string                // フィールド名
	Struct    *StructAnalysisResult // 構造体情報（再帰的）
}

func (s *StructNode) GetFieldName() string {
	return s.FieldName
}

func (s *StructNode) IsStructNode() bool {
	return true
}

func (s *StructNode) IsInterfaceNode() bool {
	return false
}

// InterfaceNode はインターフェースフィールドを表す
type InterfaceNode struct {
	FieldName      string                // フィールド名
	TypeName       string                // インターフェース型名
	PackagePath    string                // パッケージパス
	IsPointer      bool                  // ポインタ型かどうか
	ResolvedStruct *StructAnalysisResult // 解決された構造体
	Skipped        bool                  // 解決がスキップされたか
	SkipReason     string                // スキップされた理由
}

func (i *InterfaceNode) GetFieldName() string {
	return i.FieldName
}

func (i *InterfaceNode) IsStructNode() bool {
	return false
}

func (i *InterfaceNode) IsInterfaceNode() bool {
	return true
}
