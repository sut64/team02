import { TypeResearchInterface } from "./ITypeResearch";
import { AuthorNameInterface } from "./IAuthorName";
import { InstitutionNameInterface } from "./IInstitutionName";
import { MembersInterface } from "./IMember";

export interface ResearchInterface {

    ID: number,
    NameResearch: string,
    YearOfPublication: number,
    RecordingDate: Date,

    TypeResearchID: number,
    TypeResearch: TypeResearchInterface,

    AuthorNameID: number,
    AuthorName: AuthorNameInterface,

    InstitutionNameID: number,
    InstitutionName: InstitutionNameInterface,

    MemberID: number,
    Member: MembersInterface,

}