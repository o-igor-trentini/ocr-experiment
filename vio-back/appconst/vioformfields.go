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
	NumRegistro               VioFormFields = "N°"   // Registro
	Validade                  VioFormFields = "Validade"
	Observacoes               VioFormFields = "Observacoes"
	Local                     VioFormFields = "Local"
	Uf                        VioFormFields = "UF"
	DataEmissao               VioFormFields = "Data de Emissao"
	NumValidacaoCnh           VioFormFields = "Número Validacao CNH"
	NumFormRenach             VioFormFields = "Número Formulario RENACH"
)
