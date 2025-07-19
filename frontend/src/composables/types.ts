export type Category = {
    id: string;
    name: string;
    createdAt: Date;
    updatedAt: Date;
}

export type CategoryListResponse = {
    data: Category[];
    pagination: {
        limit: number;
        page: number;
        total: number;
        totalPages: number;
    }
}
