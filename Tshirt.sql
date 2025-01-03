PGDMP     3                    {           TShirt    13.11    14.8 3    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16395    TShirt    DATABASE     S   CREATE DATABASE "TShirt" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';
    DROP DATABASE "TShirt";
                postgres    false            �            1259    17107    brands    TABLE     #  CREATE TABLE public.brands (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by bigint NOT NULL,
    modified_by bigint,
    deleted_by bigint,
    name text NOT NULL
);
    DROP TABLE public.brands;
       public         heap    postgres    false            �            1259    17105    brands_id_seq    SEQUENCE     v   CREATE SEQUENCE public.brands_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.brands_id_seq;
       public          postgres    false    203            �           0    0    brands_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.brands_id_seq OWNED BY public.brands.id;
          public          postgres    false    202            �            1259    17096 
   categories    TABLE     '  CREATE TABLE public.categories (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by bigint NOT NULL,
    modified_by bigint,
    deleted_by bigint,
    name text NOT NULL
);
    DROP TABLE public.categories;
       public         heap    postgres    false            �            1259    17094    categories_id_seq    SEQUENCE     z   CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.categories_id_seq;
       public          postgres    false    201            �           0    0    categories_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;
          public          postgres    false    200            �            1259    17129    colors    TABLE     #  CREATE TABLE public.colors (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by bigint NOT NULL,
    modified_by bigint,
    deleted_by bigint,
    name text NOT NULL
);
    DROP TABLE public.colors;
       public         heap    postgres    false            �            1259    17127    colors_id_seq    SEQUENCE     v   CREATE SEQUENCE public.colors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.colors_id_seq;
       public          postgres    false    207            �           0    0    colors_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.colors_id_seq OWNED BY public.colors.id;
          public          postgres    false    206            �            1259    17140 	   materials    TABLE     &  CREATE TABLE public.materials (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by bigint NOT NULL,
    modified_by bigint,
    deleted_by bigint,
    name text NOT NULL
);
    DROP TABLE public.materials;
       public         heap    postgres    false            �            1259    17138    materials_id_seq    SEQUENCE     y   CREATE SEQUENCE public.materials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.materials_id_seq;
       public          postgres    false    209            �           0    0    materials_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.materials_id_seq OWNED BY public.materials.id;
          public          postgres    false    208            �            1259    17151    products    TABLE     �  CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by bigint NOT NULL,
    modified_by bigint,
    deleted_by bigint,
    name text NOT NULL,
    category_id bigint,
    brand_id bigint,
    size_id bigint,
    color_id bigint,
    material_id bigint,
    price numeric
);
    DROP TABLE public.products;
       public         heap    postgres    false            �            1259    17149    products_id_seq    SEQUENCE     x   CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.products_id_seq;
       public          postgres    false    211            �           0    0    products_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;
          public          postgres    false    210            �            1259    17118    sizes    TABLE     "  CREATE TABLE public.sizes (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by bigint NOT NULL,
    modified_by bigint,
    deleted_by bigint,
    name text NOT NULL
);
    DROP TABLE public.sizes;
       public         heap    postgres    false            �            1259    17116    sizes_id_seq    SEQUENCE     u   CREATE SEQUENCE public.sizes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.sizes_id_seq;
       public          postgres    false    205            �           0    0    sizes_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.sizes_id_seq OWNED BY public.sizes.id;
          public          postgres    false    204            �           2604    17110 	   brands id    DEFAULT     f   ALTER TABLE ONLY public.brands ALTER COLUMN id SET DEFAULT nextval('public.brands_id_seq'::regclass);
 8   ALTER TABLE public.brands ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202    203            �           2604    17099    categories id    DEFAULT     n   ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);
 <   ALTER TABLE public.categories ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    201    200    201            �           2604    17132 	   colors id    DEFAULT     f   ALTER TABLE ONLY public.colors ALTER COLUMN id SET DEFAULT nextval('public.colors_id_seq'::regclass);
 8   ALTER TABLE public.colors ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    207    206    207            �           2604    17143    materials id    DEFAULT     l   ALTER TABLE ONLY public.materials ALTER COLUMN id SET DEFAULT nextval('public.materials_id_seq'::regclass);
 ;   ALTER TABLE public.materials ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    208    209    209            �           2604    17154    products id    DEFAULT     j   ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);
 :   ALTER TABLE public.products ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    211    210    211            �           2604    17121    sizes id    DEFAULT     d   ALTER TABLE ONLY public.sizes ALTER COLUMN id SET DEFAULT nextval('public.sizes_id_seq'::regclass);
 7   ALTER TABLE public.sizes ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    204    205    205            �          0    17107    brands 
   TABLE DATA           t   COPY public.brands (id, created_at, modified_at, deleted_at, created_by, modified_by, deleted_by, name) FROM stdin;
    public          postgres    false    203   _;       �          0    17096 
   categories 
   TABLE DATA           x   COPY public.categories (id, created_at, modified_at, deleted_at, created_by, modified_by, deleted_by, name) FROM stdin;
    public          postgres    false    201   <       �          0    17129    colors 
   TABLE DATA           t   COPY public.colors (id, created_at, modified_at, deleted_at, created_by, modified_by, deleted_by, name) FROM stdin;
    public          postgres    false    207   �<       �          0    17140 	   materials 
   TABLE DATA           w   COPY public.materials (id, created_at, modified_at, deleted_at, created_by, modified_by, deleted_by, name) FROM stdin;
    public          postgres    false    209   �=       �          0    17151    products 
   TABLE DATA           �   COPY public.products (id, created_at, modified_at, deleted_at, created_by, modified_by, deleted_by, name, category_id, brand_id, size_id, color_id, material_id, price) FROM stdin;
    public          postgres    false    211   H>       �          0    17118    sizes 
   TABLE DATA           s   COPY public.sizes (id, created_at, modified_at, deleted_at, created_by, modified_by, deleted_by, name) FROM stdin;
    public          postgres    false    205   �?       �           0    0    brands_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.brands_id_seq', 10, true);
          public          postgres    false    202            �           0    0    categories_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.categories_id_seq', 10, true);
          public          postgres    false    200            �           0    0    colors_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.colors_id_seq', 10, true);
          public          postgres    false    206            �           0    0    materials_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.materials_id_seq', 10, true);
          public          postgres    false    208            �           0    0    products_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.products_id_seq', 1, false);
          public          postgres    false    210            �           0    0    sizes_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.sizes_id_seq', 10, true);
          public          postgres    false    204                       2606    17115    brands brands_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.brands DROP CONSTRAINT brands_pkey;
       public            postgres    false    203                       2606    17104    categories categories_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.categories DROP CONSTRAINT categories_pkey;
       public            postgres    false    201                       2606    17137    colors colors_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.colors
    ADD CONSTRAINT colors_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.colors DROP CONSTRAINT colors_pkey;
       public            postgres    false    207            	           2606    17148    materials materials_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.materials
    ADD CONSTRAINT materials_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.materials DROP CONSTRAINT materials_pkey;
       public            postgres    false    209                       2606    17159    products products_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public            postgres    false    211                       2606    17126    sizes sizes_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.sizes
    ADD CONSTRAINT sizes_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.sizes DROP CONSTRAINT sizes_pkey;
       public            postgres    false    205                       2606    17165    products fk_products_brand    FK CONSTRAINT     {   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_brand FOREIGN KEY (brand_id) REFERENCES public.brands(id);
 D   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_products_brand;
       public          postgres    false    3075    211    203                       2606    17160    products fk_products_category    FK CONSTRAINT     �   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES public.categories(id);
 G   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_products_category;
       public          postgres    false    3073    201    211                       2606    17175    products fk_products_color    FK CONSTRAINT     {   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_color FOREIGN KEY (color_id) REFERENCES public.colors(id);
 D   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_products_color;
       public          postgres    false    207    211    3079                       2606    17180    products fk_products_material    FK CONSTRAINT     �   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_material FOREIGN KEY (material_id) REFERENCES public.materials(id);
 G   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_products_material;
       public          postgres    false    209    211    3081                       2606    17170    products fk_products_size    FK CONSTRAINT     x   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_size FOREIGN KEY (size_id) REFERENCES public.sizes(id);
 C   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_products_size;
       public          postgres    false    3077    211    205            �   �   x�u�=�0����r�8i9'���*�����S�#���'=��DH�`4�6\�8nض] ���_�1�?ۚ�]R�P�v��B���F���n�uA�hq�G��V+���,)C�9����i-�,k4F��]6h+�"i���� V���H�5ɨ�y:~�@� �rGv%      �   �   x�u�M
�0�u��@$�{���M� ��&���.g�-����Ç�J�tOt�;��5AH�޿��z�t��r�A1Tg	�����dМ���K���3�i�(w���-6����n>�z��ǹ��>�6�1�{���4uc��v$�ג%��FB�~�  _�v�      �   �   x�u�M
�0�u�^ %�I=�'�&�AI�M_�3�o�1�S�����'���7:m'+���/B�&�p�;B�X�mC0J���!�YE��~���s�P��{N�aAFO��Q_B����P%]ۖ[O"'��$�G�Dm�%�q]2(���uk͏\`��su�      �   �   x�u�M
� �u�����/�C�٤���c�_]��,f��{%�r�����i;Z��B�a~�N����܎k����AM��kz�ִ�
���)��������XZh�H�=�b��˥g��8��%tʹm0q�Y�X��(gϱ�5e��� �@���@�c�x�
� _"w[      �   E  x�u�Mn�0�5y�^���p��C��8�Q��VNz�>#EV�@�D|z3$!�'����n�:��+-�?��������~l��%��Q�D3V�E9̩K����&�����7x�1��F�L�]�{�R�N��q�2>E-L�KP��MT%edՕ�<�yۯ���v~ iB�FM�:E ��p�
v�Ͽ�+txY,���a7gv��#e��7�}��FIi�$2�]�j,@|S��rCC/GSMtH-]G��\�>-���v~n�W>츴 BםpD$��ĳ�o�_�%Dd)��>h�����7�O����}���hjB���nb���Z��	�@      �   �   x�u�=�0�>H��n~8'�RH�H���ܟ"1������3����J��q�3�̋����4��_���G���Vb�F��'ː|�H�!)r���Ed\?U'!
����z/)��Q%A�Dt�8r�[� 
�L�ܡ��T0Zr�܆�Q6Xg �	r�     