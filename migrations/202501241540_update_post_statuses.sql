-- Update existing posts with 'started' status to 'in-development'
-- This migration updates the status value but the integer ID (1) remains the same
-- so existing data relationships are preserved

-- Note: The database stores status as an integer, not a string
-- The string representation is only used in the API/UI layer
-- This migration is primarily for documentation and future reference

-- No actual database changes needed since we're keeping the same integer IDs
-- and only changing the string mapping in the application layer

-- However, if you have any serialized status strings in your database
-- (like in logs or events), you may want to update those:

UPDATE events 
SET content = REPLACE(content::text, '"status":"started"', '"status":"in-development"')::jsonb
WHERE content::text LIKE '%"status":"started"%';

UPDATE logs 
SET content = REPLACE(content, '"status":"started"', '"status":"in-development"')
WHERE content LIKE '%"status":"started"%';
