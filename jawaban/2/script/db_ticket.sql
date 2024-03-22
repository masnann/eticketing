PGDMP  -                    |         	   db_ticket    16.1    16.1 C    _           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            `           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            a           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            b           1262    18384 	   db_ticket    DATABASE     �   CREATE DATABASE db_ticket WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Indonesian_Indonesia.1252';
    DROP DATABASE db_ticket;
                postgres    false            �            1259    18510    film    TABLE     �  CREATE TABLE public.film (
    id bigint NOT NULL,
    title character varying(255),
    genre character varying(255),
    year bigint,
    duration bigint,
    director character varying(255),
    description text,
    country character varying(255),
    language character varying(255),
    poster_url character varying(255),
    trailer_url character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
    DROP TABLE public.film;
       public         heap    postgres    false            �            1259    18509    film_id_seq    SEQUENCE     t   CREATE SEQUENCE public.film_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.film_id_seq;
       public          postgres    false    218            c           0    0    film_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.film_id_seq OWNED BY public.film.id;
          public          postgres    false    217            �            1259    18588    order    TABLE     #  CREATE TABLE public."order" (
    id bigint NOT NULL,
    user_id bigint,
    date timestamp with time zone,
    total_price numeric,
    order_status text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
    DROP TABLE public."order";
       public         heap    postgres    false            �            1259    18603    order_detail    TABLE     *  CREATE TABLE public.order_detail (
    id bigint NOT NULL,
    order_id bigint,
    schedule_id bigint,
    seat_id bigint,
    quantity bigint,
    price numeric,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
     DROP TABLE public.order_detail;
       public         heap    postgres    false            �            1259    18602    order_detail_id_seq    SEQUENCE     |   CREATE SEQUENCE public.order_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.order_detail_id_seq;
       public          postgres    false    228            d           0    0    order_detail_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.order_detail_id_seq OWNED BY public.order_detail.id;
          public          postgres    false    227            �            1259    18587    order_id_seq    SEQUENCE     u   CREATE SEQUENCE public.order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.order_id_seq;
       public          postgres    false    226            e           0    0    order_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.order_id_seq OWNED BY public."order".id;
          public          postgres    false    225            �            1259    18528    schedule    TABLE     2  CREATE TABLE public.schedule (
    id bigint NOT NULL,
    film_id bigint,
    studio_id bigint,
    date timestamp with time zone,
    start_time text,
    price numeric,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
    DROP TABLE public.schedule;
       public         heap    postgres    false            �            1259    18527    schedule_id_seq    SEQUENCE     x   CREATE SEQUENCE public.schedule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.schedule_id_seq;
       public          postgres    false    222            f           0    0    schedule_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.schedule_id_seq OWNED BY public.schedule.id;
          public          postgres    false    221            �            1259    18548    seat    TABLE     �   CREATE TABLE public.seat (
    id bigint NOT NULL,
    studio_id bigint,
    number text,
    status text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
    DROP TABLE public.seat;
       public         heap    postgres    false            �            1259    18547    seat_id_seq    SEQUENCE     t   CREATE SEQUENCE public.seat_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.seat_id_seq;
       public          postgres    false    224            g           0    0    seat_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.seat_id_seq OWNED BY public.seat.id;
          public          postgres    false    223            �            1259    18520    studio    TABLE     �   CREATE TABLE public.studio (
    id bigint NOT NULL,
    name character varying(255),
    capacity bigint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
    DROP TABLE public.studio;
       public         heap    postgres    false            �            1259    18519    studio_id_seq    SEQUENCE     v   CREATE SEQUENCE public.studio_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.studio_id_seq;
       public          postgres    false    220            h           0    0    studio_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.studio_id_seq OWNED BY public.studio.id;
          public          postgres    false    219            �            1259    18396    users    TABLE     �  CREATE TABLE public.users (
    id bigint NOT NULL,
    email character varying(255),
    password character varying(255),
    phone character varying(255),
    role character varying(255),
    name character varying(255),
    photo_profile character varying(255),
    gender character varying(255),
    date_of_birth date,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    18395    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    216            i           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    215            �           2604    18513    film id    DEFAULT     b   ALTER TABLE ONLY public.film ALTER COLUMN id SET DEFAULT nextval('public.film_id_seq'::regclass);
 6   ALTER TABLE public.film ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            �           2604    18591    order id    DEFAULT     f   ALTER TABLE ONLY public."order" ALTER COLUMN id SET DEFAULT nextval('public.order_id_seq'::regclass);
 9   ALTER TABLE public."order" ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    225    226    226            �           2604    18606    order_detail id    DEFAULT     r   ALTER TABLE ONLY public.order_detail ALTER COLUMN id SET DEFAULT nextval('public.order_detail_id_seq'::regclass);
 >   ALTER TABLE public.order_detail ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    227    228    228            �           2604    18531    schedule id    DEFAULT     j   ALTER TABLE ONLY public.schedule ALTER COLUMN id SET DEFAULT nextval('public.schedule_id_seq'::regclass);
 :   ALTER TABLE public.schedule ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    222    222            �           2604    18551    seat id    DEFAULT     b   ALTER TABLE ONLY public.seat ALTER COLUMN id SET DEFAULT nextval('public.seat_id_seq'::regclass);
 6   ALTER TABLE public.seat ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    224    223    224            �           2604    18523 	   studio id    DEFAULT     f   ALTER TABLE ONLY public.studio ALTER COLUMN id SET DEFAULT nextval('public.studio_id_seq'::regclass);
 8   ALTER TABLE public.studio ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    219    220            �           2604    18399    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            R          0    18510    film 
   TABLE DATA           �   COPY public.film (id, title, genre, year, duration, director, description, country, language, poster_url, trailer_url, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    218   �M       Z          0    18588    order 
   TABLE DATA           s   COPY public."order" (id, user_id, date, total_price, order_status, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    226   �O       \          0    18603    order_detail 
   TABLE DATA              COPY public.order_detail (id, order_id, schedule_id, seat_id, quantity, price, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    228   P       V          0    18528    schedule 
   TABLE DATA           w   COPY public.schedule (id, film_id, studio_id, date, start_time, price, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    222   7P       X          0    18548    seat 
   TABLE DATA           a   COPY public.seat (id, studio_id, number, status, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    224   �P       T          0    18520    studio 
   TABLE DATA           X   COPY public.studio (id, name, capacity, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    220   �P       P          0    18396    users 
   TABLE DATA           �   COPY public.users (id, email, password, phone, role, name, photo_profile, gender, date_of_birth, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    216   >Q       j           0    0    film_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.film_id_seq', 3, true);
          public          postgres    false    217            k           0    0    order_detail_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.order_detail_id_seq', 1, false);
          public          postgres    false    227            l           0    0    order_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.order_id_seq', 1, false);
          public          postgres    false    225            m           0    0    schedule_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.schedule_id_seq', 4, true);
          public          postgres    false    221            n           0    0    seat_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.seat_id_seq', 1, false);
          public          postgres    false    223            o           0    0    studio_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.studio_id_seq', 3, true);
          public          postgres    false    219            p           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 3, true);
          public          postgres    false    215            �           2606    18517    film film_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.film
    ADD CONSTRAINT film_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.film DROP CONSTRAINT film_pkey;
       public            postgres    false    218            �           2606    18610    order_detail order_detail_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.order_detail
    ADD CONSTRAINT order_detail_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.order_detail DROP CONSTRAINT order_detail_pkey;
       public            postgres    false    228            �           2606    18595    order order_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public."order" DROP CONSTRAINT order_pkey;
       public            postgres    false    226            �           2606    18535    schedule schedule_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.schedule
    ADD CONSTRAINT schedule_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.schedule DROP CONSTRAINT schedule_pkey;
       public            postgres    false    222            �           2606    18555    seat seat_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.seat
    ADD CONSTRAINT seat_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.seat DROP CONSTRAINT seat_pkey;
       public            postgres    false    224            �           2606    18525    studio studio_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.studio
    ADD CONSTRAINT studio_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.studio DROP CONSTRAINT studio_pkey;
       public            postgres    false    220            �           2606    18403    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    216            �           1259    18518    idx_film_deleted_at    INDEX     J   CREATE INDEX idx_film_deleted_at ON public.film USING btree (deleted_at);
 '   DROP INDEX public.idx_film_deleted_at;
       public            postgres    false    218            �           1259    18601    idx_order_deleted_at    INDEX     N   CREATE INDEX idx_order_deleted_at ON public."order" USING btree (deleted_at);
 (   DROP INDEX public.idx_order_deleted_at;
       public            postgres    false    226            �           1259    18626    idx_order_detail_deleted_at    INDEX     Z   CREATE INDEX idx_order_detail_deleted_at ON public.order_detail USING btree (deleted_at);
 /   DROP INDEX public.idx_order_detail_deleted_at;
       public            postgres    false    228            �           1259    18546    idx_schedule_deleted_at    INDEX     R   CREATE INDEX idx_schedule_deleted_at ON public.schedule USING btree (deleted_at);
 +   DROP INDEX public.idx_schedule_deleted_at;
       public            postgres    false    222            �           1259    18561    idx_seat_deleted_at    INDEX     J   CREATE INDEX idx_seat_deleted_at ON public.seat USING btree (deleted_at);
 '   DROP INDEX public.idx_seat_deleted_at;
       public            postgres    false    224            �           1259    18526    idx_studio_deleted_at    INDEX     N   CREATE INDEX idx_studio_deleted_at ON public.studio USING btree (deleted_at);
 )   DROP INDEX public.idx_studio_deleted_at;
       public            postgres    false    220            �           1259    18404    idx_users_deleted_at    INDEX     L   CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
 (   DROP INDEX public.idx_users_deleted_at;
       public            postgres    false    216            �           2606    18616 %   order_detail fk_order_detail_schedule    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_detail
    ADD CONSTRAINT fk_order_detail_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedule(id);
 O   ALTER TABLE ONLY public.order_detail DROP CONSTRAINT fk_order_detail_schedule;
       public          postgres    false    222    4783    228            �           2606    18621 !   order_detail fk_order_detail_seat    FK CONSTRAINT        ALTER TABLE ONLY public.order_detail
    ADD CONSTRAINT fk_order_detail_seat FOREIGN KEY (seat_id) REFERENCES public.seat(id);
 K   ALTER TABLE ONLY public.order_detail DROP CONSTRAINT fk_order_detail_seat;
       public          postgres    false    228    224    4786            �           2606    18611 "   order_detail fk_order_order_detail    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_detail
    ADD CONSTRAINT fk_order_order_detail FOREIGN KEY (order_id) REFERENCES public."order"(id);
 L   ALTER TABLE ONLY public.order_detail DROP CONSTRAINT fk_order_order_detail;
       public          postgres    false    4789    226    228            �           2606    18596    order fk_order_user    FK CONSTRAINT     t   ALTER TABLE ONLY public."order"
    ADD CONSTRAINT fk_order_user FOREIGN KEY (user_id) REFERENCES public.users(id);
 ?   ALTER TABLE ONLY public."order" DROP CONSTRAINT fk_order_user;
       public          postgres    false    226    216    4774            �           2606    18536    schedule fk_schedule_film    FK CONSTRAINT     w   ALTER TABLE ONLY public.schedule
    ADD CONSTRAINT fk_schedule_film FOREIGN KEY (film_id) REFERENCES public.film(id);
 C   ALTER TABLE ONLY public.schedule DROP CONSTRAINT fk_schedule_film;
       public          postgres    false    218    4776    222            �           2606    18541    schedule fk_schedule_studio    FK CONSTRAINT     }   ALTER TABLE ONLY public.schedule
    ADD CONSTRAINT fk_schedule_studio FOREIGN KEY (studio_id) REFERENCES public.studio(id);
 E   ALTER TABLE ONLY public.schedule DROP CONSTRAINT fk_schedule_studio;
       public          postgres    false    222    220    4780            �           2606    18556    seat fk_seat_studio    FK CONSTRAINT     u   ALTER TABLE ONLY public.seat
    ADD CONSTRAINT fk_seat_studio FOREIGN KEY (studio_id) REFERENCES public.studio(id);
 =   ALTER TABLE ONLY public.seat DROP CONSTRAINT fk_seat_studio;
       public          postgres    false    224    220    4780            R   �  x����n�0���S��qҴ�-K�a��a�'3�j���8�ۏr�.��CO������Y��
�I9[lD^ưiz�ihwBMު��fU1[\[TL�K
p�4r$�h� &BA��]�DIJ�%��[�+�>�4��L�Ġl���N��*B���A�lO��.Gy�2� ��B$7(��MV l˛�SY|��7��*�B���z:�G4^S)��z�冩��~��۳�P�?��/I��T�I]Cu��X��y�\͗�˿�ߎ��3�{����V�2=s�e��u@����*f��_�O:�,
�κ��ÿ�c��v$�N0��D�=�yr�\���%�fo0Fj�1��'kEַ�-�x2ۤ�!���8l�y����J�t�f�[d�8��b_�� �әS�����4��\����AL9��E1[���?���v���ٓ~����9���g���cJd�=s#�b�{��QF�m�z������}"~��r4� ��l�      Z      x������ � �      \      x������ � �      V   }   x���11k��C���~D^p�Gl])�,���F���]��bØ��5s�c��E�?�Jكm�R���S�L�.�2w(d���j��`;����:��)ue�?��|�n��~,�Pi      X      x������ � �      T   M   x�3�.)M��W0�440�4202�50�52R00�25�20�3��017�'��e3ǈ�Јs�a�sZP`L� ��,�      P   1  x�}��n�@���)|p�r�*hѪ�-�oЂ�rX���w�m��d���&�@��?�<H2-�9(���x!L 4e�)3Lܛ�������˚_�,���
��[�f�!T!�%G����M0��a]���N��Ln�^e�7Q��,�K����mϮ��=��x���YO���ri�o[��E�	����e �n�ۘؔiHG�Yc��������gw�M�ό����q.����d�H���j�{�����5�nS���+��m��	��(O��{�����L
-6I�IS���     