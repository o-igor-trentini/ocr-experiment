package appconst

type VioFormFields string

const (
	Nome                      VioFormFields = "Nome"
	DocIdentidadeOrgEmissorUf VioFormFields = "Doc." // Doc Identidade/Org. Emissor/UF
	Cpf                       VioFormFields = "CPF"
	DataDeNascimento          VioFormFields = "Data" // Data de nascimento
	FiliacaoPai               VioFormFields = "Pai"  // Filiacao Pai
	FiliacaoMae               VioFormFields = "Mae"  // Filiacao Mãe
	Permissao                 VioFormFields = "Permissao"
	Acc                       VioFormFields = "ACC"
	CatHabilitacao            VioFormFields = "Cat." // Hab.
	NumRegistro               VioFormFields = "N°"   // N° Registro
	Validade                  VioFormFields = "Validade"
	PrimHabilitacao           VioFormFields = "Habilita" // 1a Habilitacao
	Observacoes               VioFormFields = "Observa"  // Observacoes
	Local                     VioFormFields = "Local"
	Uf                        VioFormFields = "UF"
	DataEmissao               VioFormFields = "Emissao" // Data de Emissao
	NumValidacaoCnh           VioFormFields = "CNH"     // Numero Validacao CNH
	NumFormRenach             VioFormFields = "RENACH"  // Numero Formulario RENACH
)
