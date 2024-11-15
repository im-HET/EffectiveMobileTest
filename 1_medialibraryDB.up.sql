--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2024-11-13 15:12:48

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4850 (class 1262 OID 24577)
-- Name: mediaLibraryDB; Type: DATABASE; Schema: -; Owner: mediaLibraryUser
--

CREATE DATABASE "mediaLibraryDB2" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';


ALTER DATABASE "mediaLibraryDB2" OWNER TO "mediaLibraryUser";

\connect "mediaLibraryDB"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 219 (class 1255 OID 32793)
-- Name: getverse(integer, integer); Type: FUNCTION; Schema: public; Owner: mediaLibraryUser
--

CREATE FUNCTION public.getverse(i integer, n integer) RETURNS TABLE(amount integer, num integer, verse text)
    LANGUAGE plpgsql STRICT
    AS $$
	DECLARE
	arr TEXT ARRAY;
	tx TEXT;
BEGIN
	select library.text INTO tx from library where library.id = i;
	arr:= regexp_split_to_array(tx, '\n\n');
	amount:= array_length(arr, 1);
	IF amount > 0 THEN
		IF amount >= n THEN
			num = n;
			verse = arr[n];
		ELSE 
			num = 0;
			verse = '';
		END IF;
	END IF;
	RETURN NEXT;
END;
$$;


ALTER FUNCTION public.getverse(i integer, n integer) OWNER TO "mediaLibraryUser";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 24586)
-- Name: groups; Type: TABLE; Schema: public; Owner: mediaLibraryUser
--

CREATE TABLE public.groups (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.groups OWNER TO "mediaLibraryUser";

--
-- TOC entry 218 (class 1259 OID 24593)
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: mediaLibraryUser
--

ALTER TABLE public.groups ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
    CYCLE
);


--
-- TOC entry 215 (class 1259 OID 24578)
-- Name: library; Type: TABLE; Schema: public; Owner: mediaLibraryUser
--

CREATE TABLE public.library (
    id integer NOT NULL,
    song text,
    releasedate date,
    text text,
    link text,
    groupid integer
);


ALTER TABLE public.library OWNER TO "mediaLibraryUser";

--
-- TOC entry 216 (class 1259 OID 24585)
-- Name: mediaLibraryMainTable_id_seq; Type: SEQUENCE; Schema: public; Owner: mediaLibraryUser
--

ALTER TABLE public.library ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public."mediaLibraryMainTable_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
    CYCLE
);


--
-- TOC entry 4843 (class 0 OID 24586)
-- Dependencies: 217
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: mediaLibraryUser
--

INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (1, 'Кино') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (2, 'Пикник') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (3, 'Moby') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (4, 'Five Finger Death Punch') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (5, 'The Hu') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (6, 'Nickelback') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (7, 'Nautilus Pompilius') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (8, 'Foo Fighters') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (9, 'Pink Floyd') ON CONFLICT DO NOTHING;
INSERT INTO public.groups OVERRIDING SYSTEM VALUE VALUES (10, 'Deep Purple
') ON CONFLICT DO NOTHING;


--
-- TOC entry 4841 (class 0 OID 24578)
-- Dependencies: 215
-- Data for Name: library; Type: TABLE DATA; Schema: public; Owner: mediaLibraryUser
--

INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (1, 'Место для шага вперед', '1989-01-01', 'У меня есть дом, только нет ключей,
 у меня есть солнце, но оно среди туч,
Есть голова, только нет плечей,
 но я вижу, как тучи режут солнечный луч.
У меня есть слово, но в нем нет букв,
 у меня есть лес, но нет топоров,
У меня есть время, но нет сил ждать,
 и есть еще ночь, но в ней нет снов.

И есть еще белые, белые дни,
 белые горы и белый лед.
Но все, что мне нужно - это несколько слов
 и место для шага вперед.

У меня река, только нет моста,
 у меня есть мыши, но нет кота,
У меня есть парус, но ветра нет
 и есть еще краски, но нет холста.
У меня на кухне из крана вода,
 у меня есть рана, но нет бинта,
У меня есть братья, но нет родных
 и есть рука, и она пуста.

И есть еще белые, белые дни,
 белые горы и белый лед.
Но все, что мне нужно - это несколько слов
 и место для шага вперед.

И есть еще белые, белые дни,
 белые горы и белый лед.
Но все, что мне нужно - это несколько слов
 и место для шага вперед.', 'https://youtube.com/', 1) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (4, 'I Apologize', '2015-09-02', 'One day the shadows will surround me
Someday the days will come to end
Sometime I''ll have to face the real me
Somehow I''ll have to learn to bend
And now I see clearly
All these times I simply stepped aside
I watched but never really listened
As the whole world passed me by
All this time I watched from the outside
Never understood what was wrong or what was right
I apologize, whoah
I apologize, whoah
One day I''ll face the Hell inside me
Someday I''ll accept what I have done
Sometime I''ll leave the past behind me
For now I accept who I''ve become
And now I see clearly
All these times I simply stepped aside
I watched but never really listened
As the whole world passed me by
All this time I watched from the outside
Never understood what was wrong or what was right
I apologize
I apologize
One day the shadows will surround me
All these times I simply stepped aside
I watched but never really listened
As the whole world passed me by
All this time I watched from the outside
Never understood what was wrong or what was right
I apologize
I apologize
Whoah, I apologize
Whoah, I apologize
Whoah, I apologize
Whoah, I apologize
I apologize', NULL, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (5, 'Wolf Totem', NULL, NULL, 'https://youtube.com/', 5) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (10, 'Another Brick in the Wall', '1979-11-30', NULL, 'https://youtube.com/', 9) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (2, 'Королевство кривых', '2005-04-01', 'Огнями реклам,
Неоновых ламп
Бьет город мне в спину, торопит меня.
А я не спешу,
Я этим дышу
И то, что мое, ему не отнять.

Минуту еще, мой ветер не стих,
Мне нравится здесь в Королевстве Кривых.
Минуту еще, мой ветер не стих,
Мне нравится здесь в Королевстве Кривых.

Здесь деньги не ждут,
Когда их сожгут
В их власти, дать счастье и счастье отнять.
Но только не мне,
Я сам по себе
И темные улицы манят меня.

Минуту еще, мой ветер не стих
Мне нравится здесь ...
Минуту еще, мой ветер не стих
Мне нравится здесь ...

Он занят игрой
И каждый второй,
Да каждый второй замедляет свой шаг.
Но только не я,
Я весел и пьян,
Я только сейчас начинаю дышать.

Минуту еще...
Минуту еще...', 'http://youtube.com', 2) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (3, 'Natural Blues', '2000-03-06', 'Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Went down the hill, the other day
Soul got happy and stayed all day
Went down the hill, the other day
Soul got happy and stayed all day
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Went in the room, didn''t stay long
Looked on the bed and brother was dead
Went in the room, didn''t stay long
Looked on the bed and brother was dead
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God
Oh Lordy, Lord, trouble so hard
Oh Lordy, Lord, trouble so hard
Don''t nobody know my troubles but God
Don''t nobody know my troubles but God', 'http://youtube.com', 3) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (6, 'Rockstar', NULL, 'I''m through with standing in line to clubs I''ll never get in
It''s like the bottom of the ninth, and I''m never gonna win
This life hasn''t turned out quite the way I want it to be
(Tell me what you want)
I want a brand new house on an episode of Cribs
And a bathroom I can play baseball in
And a king-size tub, big enough for ten plus me
(Uh, so what you need?)
I''ll need a credit card that''s got no limit
And a big black jet with a bedroom in it
Gonna join the mile high club at 37, 000 feet
(Been there, done that)
I want a new tour bus full of old guitars
My own star on Hollywood Boulevard
Somewhere between Cher and James Dean is fine for me
(So how you gonna do it?)
I''m gonna trade this life for fortune and fame
I''d even cut my hair and changed my name
''Cause we all just wanna be big rockstars
And live in hilltop houses, driving 15 cars
The girls come easy, and the drugs come cheap
We''ll all stay skinny, ''cause we just won''t eat
And we''ll hang out in the coolest bars
In the V.I.P. with the movie stars
Every good gold digger''s gonna wind up there
Every Playboy Bunny with her bleached blond hair and, well
Hey, hey, I wanna be a rockstar
Hmm, hey, hey, I wanna be a rockstar
I wanna be great like Elvis without the tassels
Hire eight body guards that love to beat up assholes
Sign a couple autographs, so I can eat my meals for free
(I''ll have the quesadilla, haha!)
I''m gonna dress my ass with the latest fashion
Get a front door key to the Playboy mansion
Gonna date a centerfold that loves to blow my money for me
(So how you gonna do it?)
I''m gonna trade this life for fortune and fame
I''d even cut my hair and changed my name
''Cause we all just wanna be big rockstars
And live in hilltop houses, driving 15 cars
The girls come easy, and the drugs come cheap
We''ll all stay skinny, ''cause we just won''t eat
And we''ll hang out in the coolest bars
In the V.I.P. with the movie stars
Every good gold digger''s gonna wind up there
Every Playboy Bunny with her bleached blond hair
And we''ll hide out in the private rooms
With the latest dictionary, and today''s who''s who
They''ll get you anything with that evil smile
Everybody''s got a drug dealer on speed dial, well
Hey, hey, I wanna be a rockstar
I''m gonna sing those songs that offend the censors
Gonna pop my pills from a Pez dispenser
I''ll get washed-up singers writing all my songs
Lip sync ''em every night, so I don''t get ''em wrong
Well, we all just wanna be big rockstars
And live in hilltop houses, driving 15 cars
The girls come easy, and the drugs come cheap
We''ll all stay skinny, ''cause we just won''t eat
And we''ll hang out in the coolest bars
In the V.I.P. with the movie stars
Every good gold digger''s gonna wind up there
Every Playboy Bunny with her bleached blond hair
And we''ll hide out in the private rooms
With the latest dictionary, and today''s who''s who
They''ll get you anything with that evil smile
Everybody''s got a drug dealer on speed dial, well
Hey, hey, I wanna be a rockstar
Hmm, hey, hey, I wanna be a rockstar', NULL, 6) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (7, 'Lift Me Up', '2005-02-28', 'Plain talking (plain talking)
Take us so far (take us so far)
Broken down cars (broken down cars)
Like strung out old stars (like strung out old stars)
Plain talking (plain talking)
Served us so well (served us so well)
Traveled through hell (traveled through hell)
And oh, how we fell (oh, how we fell)
Lift me up, lift me up
Higher now, Ama
Push me up, lift me up
Higher now, Ama
Push me up, lift me up
Higher now, Ama
Lift me up, lift me up
Higher now, Ama
Plain talking (plain talking)
Making us bold (making us bold)
So strung out and cold (so strung out and cold)
I''m feeling so old (feeling so old)
Plain talking (plain talking)
Has ruined us now (has ruined us now)
You never know how (you never know how)
Sweeter than thou (sweeter than thou)
Lift me up, lift me up
Higher now, Ama
Lift me up, lift me up
Higher now, Ama
Push me up, lift me up
Higher, now Ama
Lift me up, lift me up
Higher now, Ama
Lift me up, lift me up
Higher now, Ama
Push me up, lift me up
Higher now, Ama
Lift me up, lift me up
Higher now, Ama
Lift me up, lift me up
Higher, now Ama
Lift me up, lift me up (oh, la, la, la, la)
Lift me up, lift me up (oh, la, la, la, la)
Lift me up, lift me up (oh, la, la, la, la)
(Feeling so bored)
Lift me up, lift me up (oh, la, la, la, la)
(Feeling so bored)
Lift me up, lift me up (oh, la, la, la, la)
(Feeling so bored)
Lift me up, lift me up (oh, la, la, la, la)
Lift me up, lift me up (oh, la, la, la, la)
(Feeling so bored)
Lift me up, lift me up (oh, la, la, la, la)', NULL, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (8, 'Бриллиантовые дороги', '1988-08-04', 'Посмотри, как блестят
Бриллиантовые дороги.
Послушай, как хрустят
Бриллиантовые дороги.
Смотри, какие следы
Оставляют на них Боги.
Чтоб идти вслед за ними нужны
Золотые ноги.
Чтоб вцепиться в стекло,
Нужны алмазные когти.
Горят над нами, горят,
Помрачая рассудок,
Бриллиантовые дороги
В темное время суток.
.
Посмотри, как узки
Бриллиантовые дороги.
Нас зажали в тиски
Бриллиантовые дороги.
Чтобы видеть их свет
Мы пили горькие травы.
Если в пропасть не пасть,
Все равно умирать от отравы,
На алмазных мостах,
Через черные канавы.
Парят над нами, парят,
Помрачая рассудок,
Бриллиантовые дороги
В темное время суток.', NULL, 7) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (9, 'The Pretender', '2007-08-21', 'Keep you in the dark
You know they all pretend
Keep you in the dark
And so it all began
Send in your skeletons
Sing as their bones come marchin'' in again
The need you buried deep
The secrets that you keep are at the ready
Are you ready?
I''m finished making sense
Done pleading ignorance, that whole defense
Spinning infinity, but
The wheel is spinning me, it''s never-ending, never-ending
Same old story
What if I say I''m not like the others?
What if I say I''m not just another one of your plays?
You''re the pretender
What if I say I will never surrender?
What if I say I''m not like the others?
What if I say I''m not just another one of your plays?
You''re the pretender
What if I say I will never surrender?
In time, or so I''m told
I''m just another soul for sale, oh well
The page is out of print
We are not permanent, we''re temporary, temporary
Same old story
What if I say I''m not like the others?
What if I say I''m not just another one of your plays?
You''re the pretender
What if I say I will never surrender?
What if I say I''m not like the others?
What if I say I''m not just another one of your plays?
You''re the pretender
What if I say I will never surrender? Oh
I''m the voice inside your head you refuse to hear
I''m the face that you have to face, mirroring your stare
I''m what''s left, I''m what''s right, I''m the enemy
I''m the hand that''ll take you down, bring you to your knees
So, who are you?
Yeah, who are you?
Yeah, who are you?
Yeah, who are you?
Keep you in the dark
You know they all pretend
What if I say I''m not like the others?
What if I say I''m not just another one of your plays?
You''re the pretender
What if I say I will never surrender?
What if I say I''m not like the others?
What if I say I''m not just another one of your plays?
You''re the pretender
What if I say I will never surrender?
What if I say I''m not like the others? (Keep you in the dark)
What if I say I''m not just another one of your plays? (You know they all)
You''re the pretender (pretend)
What if I say I will never surrender?
What if I say I''m not like the others? (Keep you in the dark)
What if I say I''m not just another one of your plays? (You know they all)
You''re the pretender (pretend)
What if I say I will never surrender?
So, who are you?
Yeah, who are you?
Yeah, who are you?', NULL, 8) ON CONFLICT DO NOTHING;
INSERT INTO public.library OVERRIDING SYSTEM VALUE VALUES (11, 'Smoke on the Water', NULL, 'We all came out to Montreux
On the Lake Geneva shoreline
To make records with a mobile, yeah
We didn''t have much time now
Frank Zappa and the Mothers
Were at the best place around
But some stupid with a flare gun
Burned the place to the ground
Smoke on the water, a fire in the sky
(Smoke) on the water, you guys are great
They burned down the gambling house
It died with an awful sound
Funky Claude was running in and out
He was pulling kids out the ground now
When it all was over
Find another place
Swiss time was running out
It seemed that we would lose the race
Smoke on the water, a fire in the sky
Smoke on the water
Burn it down
We ended up at the Grand Hotel
It was empty, cold and bare
The Rolling truck Stones thing just outside
Huh, making our music there now
With a few red lights and a few old beds
We made a place to sweat
No matter what we get out of this
I know, I know we''ll never forget
Smoke on the water, a fire in the sky
Smoke on the water
(I can''t hear anything)
one more time
(Smoke on the water) hey!', NULL, 10) ON CONFLICT DO NOTHING;


--
-- TOC entry 4851 (class 0 OID 0)
-- Dependencies: 218
-- Name: groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: mediaLibraryUser
--

SELECT pg_catalog.setval('public.groups_id_seq', 15, true);


--
-- TOC entry 4852 (class 0 OID 0)
-- Dependencies: 216
-- Name: mediaLibraryMainTable_id_seq; Type: SEQUENCE SET; Schema: public; Owner: mediaLibraryUser
--

SELECT pg_catalog.setval('public."mediaLibraryMainTable_id_seq"', 15, true);


--
-- TOC entry 4697 (class 2606 OID 24592)
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: mediaLibraryUser
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- TOC entry 4695 (class 2606 OID 24584)
-- Name: library mediaLibraryMainTable_pkey; Type: CONSTRAINT; Schema: public; Owner: mediaLibraryUser
--

ALTER TABLE ONLY public.library
    ADD CONSTRAINT "mediaLibraryMainTable_pkey" PRIMARY KEY (id);


-- Completed on 2024-11-13 15:12:48

--
-- PostgreSQL database dump complete
--

